package repository

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"main-server/pkg/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// Константы для конфигурирования
const (
	tokenTTL_access  = 1 * time.Hour
	tokenTTL_refresh = 12 * time.Hour
)

type AuthPostgres struct {
	db *sqlx.DB
}

/*
* Функция создания экземпляра сервиса
 */
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

/*
* Функция регистрации пользователя
 */
func (r *AuthPostgres) CreateUser(user model.UserRegisterModel) (model.UserAuthDataModel, error) {
	check := CheckRowExists(r.db, usersTable, "email", user.Email)

	if check {
		return model.UserAuthDataModel{}, errors.New("Пользователь с данным email-адресом уже существует!")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	// Хэширование пароля
	// user.Password = generatePasswordHash(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), viper.GetInt("crypt.cost"))
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	user.Password = string(hashedPassword)

	var id int
	var userUuid string
	query := fmt.Sprintf("INSERT INTO %s (email, password, uuid) values ($1, $2, $3) RETURNING id, uuid", usersTable)

	// Generate UUID
	u1 := uuid.NewV4()

	row := tx.QueryRow(query, user.Email, user.Password, u1)
	if err := row.Scan(&id, &userUuid); err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, errors.New("Пользователь с данными регистрационными данными уже существует!")
	}

	// Добавление пользовательских данных
	query = fmt.Sprintf("INSERT INTO %s (name, surname, date_registration, users_id) values ($1, $2, $3, $4)", usersDataTable)
	_, err = tx.Exec(query, user.Name, user.Surname, time.Now(), id)
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE value = $1 limit 1", rolesTable)
	var role model.RoleModel
	err = r.db.Get(&role, query, "USER")
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, errors.New("Роли пользователя не существует в базе данных!")
	}

	// Добавление роли пользователю (по-умолчанию данная роль - USER)
	query = fmt.Sprintf("INSERT INTO %s (users_id, roles_id) VALUES ($1, $2)", usersRolesTable)
	_, err = tx.Exec(query, id, role.Id)
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	// Генерация токенов доступа и обновления
	accessToken, err := GenerateToken(userUuid, role.Uuid, tokenTTL_access, viper.GetString("token.signing_key_access"))
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	refreshToken, err := GenerateToken(userUuid, role.Uuid, tokenTTL_refresh, viper.GetString("token.signing_key_refresh"))
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	// Установка токенов пользователю
	query = fmt.Sprintf("INSERT INTO %s (users_id, access_token, refresh_token) values ($1, $2, $3)", tokensTable)
	_, err = tx.Exec(query, id, accessToken, refreshToken)
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	return model.UserAuthDataModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, tx.Commit()
}

/*
* Функция авторизации пользователя
 */
func (r *AuthPostgres) LoginUser(user model.UserLoginModel) (model.UserAuthDataModel, error) {
	var findUser model.UserModel
	query := fmt.Sprintf("SELECT * FROM %s tl WHERE tl.email = $1", usersTable)
	if err := r.db.Get(&findUser, query, user.Email); err != nil {
		return model.UserAuthDataModel{}, errors.New("Пользователя с данным почтовым адресом не существует!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password)); err != nil {
		return model.UserAuthDataModel{}, errors.New("Не правильный пароль! Повторите попытку")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	query = fmt.Sprintf("DELETE FROM %s tl WHERE tl.users_id = $1", tokensTable)
	if _, err := r.db.Exec(query, findUser.Id); err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	query = fmt.Sprintf(`SELECT roles.id, roles.uuid, roles.value, roles.description, roles.users_id FROM %s 
			INNER JOIN %s on users_roles.roles_id = roles.id WHERE users_roles.users_id = $1`, usersRolesTable, rolesTable)

	var role model.RoleModel
	if err := r.db.Get(&role, query, findUser.Id); err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, errors.New("Пользователь не имеет роли!")
	}

	// Генерация токенов доступа и обновления
	accessToken, err := GenerateToken(findUser.Uuid, role.Uuid, tokenTTL_access, viper.GetString("token.signing_key_access"))
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	refreshToken, err := GenerateToken(findUser.Uuid, role.Uuid, tokenTTL_refresh, viper.GetString("token.signing_key_refresh"))
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	// Установка токенов пользователю
	query = fmt.Sprintf("INSERT INTO %s (users_id, access_token, refresh_token) values ($1, $2, $3)", tokensTable)
	_, err = tx.Exec(query, findUser.Id, accessToken, refreshToken)
	if err != nil {
		tx.Rollback()
		return model.UserAuthDataModel{}, err
	}

	return model.UserAuthDataModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, tx.Commit()
}

/*
* Функция обновления токена доступа
 */
func (r *AuthPostgres) Refresh(token model.TokenRefreshModel) (model.UserAuthDataModel, error) {
	userData, err := ParseTokenWithoutValid(token.RefreshToken, viper.GetString("token.signing_key_refresh"))
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	user, err := r.GetUser("uuid", userData.UsersId)
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	var findToken model.TokenModel
	query := fmt.Sprintf("SELECT * FROM %s tl WHERE tl.refresh_token = $1 AND tl.users_id = $2", tokensTable)

	if err := r.db.Get(&findToken, query, token.RefreshToken, user.Id); err != nil {
		return model.UserAuthDataModel{}, errors.New("Пользователя с данным токеном обновления не существует!")
	}

	query = fmt.Sprintf(`SELECT roles.id, roles.uuid, roles.value, roles.description, roles.users_id FROM %s 
			INNER JOIN %s on users_roles.roles_id = roles.id WHERE users_roles.users_id = $1`, usersRolesTable, rolesTable)

	var role model.RoleModel
	if err := r.db.Get(&role, query, user.Id); err != nil {
		return model.UserAuthDataModel{}, errors.New("Пользователь не имеет роли!")
	}

	isValid := ValidToken(token.RefreshToken, viper.GetString("token.signing_key_refresh"))

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	var refreshToken string

	logrus.Info(isValid)
	if !isValid {
		refreshToken, err = GenerateToken(user.Uuid, role.Uuid, tokenTTL_refresh, viper.GetString("token.signing_key_refresh"))
		if err != nil {
			return model.UserAuthDataModel{}, err
		}

		setValues = append(setValues, fmt.Sprintf("refresh_token=$%d", argId))
		args = append(args, refreshToken)
		argId++
	} else {
		refreshToken = token.RefreshToken
	}

	accessToken, err := GenerateToken(user.Uuid, role.Uuid, tokenTTL_access, viper.GetString("token.signing_key_access"))
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	setValues = append(setValues, fmt.Sprintf("access_token=$%d", argId))
	args = append(args, accessToken)
	argId++

	setQuery := strings.Join(setValues, ", ")

	query = fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.users_id = $%d",
		tokensTable, setQuery, argId)
	args = append(args, user.Id)

	// Обновление данных о токене пользователя
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return model.UserAuthDataModel{}, err
	}

	// Возвращение авторизационных данных пользователя
	return model.UserAuthDataModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

/*
* Функция разлогирования пользователя
 */
func (r *AuthPostgres) Logout(tokens model.TokenDataModel) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.access_token=$1 AND tl.refresh_token=$2 RETURNING id", tokensTable)
	row := r.db.QueryRow(query, tokens.AccessToken, tokens.RefreshToken)

	var id int
	if err := row.Scan(&id); err != nil {
		return false, err
	}

	// logrus.Info(string(rune(id)))

	return true, nil
}

/*
* Функция получения данных о пользователе
 */
func (r *AuthPostgres) GetUser(column, value string) (model.UserModel, error) {
	var user model.UserModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", usersTable, column)

	err := r.db.Get(&user, query, value)

	return user, err
}

/*
* Функция получения данных о роли
 */
func (r *AuthPostgres) GetRole(column, value string) (model.RoleModel, error) {
	var user model.RoleModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", rolesTable, column)

	err := r.db.Get(&user, query, value)

	return user, err
}

/*
* Функция хэширования пароля
 */
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("crypt.salt"))))
}

// Структура определяющая данные токена
type tokenClaims struct {
	jwt.StandardClaims
	UsersId string `json:"users_id"`
	RolesId string `json:"roles_id"`
}

/*
* Функция генерации токена
 */
func GenerateToken(uuid, rolesUuid string, tokenTTL time.Duration, signingKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		uuid,
		rolesUuid,
	})

	return token.SignedString([]byte(signingKey))
}

/*
* Функция получения данных из токена без проверки на валидацию
 */
func ParseTokenWithoutValid(pToken, signingKey string) (model.TokenOutputParseString, error) {
	token, _ := jwt.ParseWithClaims(pToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	// Получение данных из токена (с преобразованием к указателю на tokenClaims)
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return model.TokenOutputParseString{}, errors.New("token claims are not of type")
	}

	return model.TokenOutputParseString{
		UsersId: claims.UsersId,
		RolesId: claims.RolesId,
	}, nil
}

/*
* Функция проверки валидности токена
 */
func ValidToken(pToken, signingKey string) bool {
	_, err := jwt.ParseWithClaims(pToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return false
	}

	return true
}
