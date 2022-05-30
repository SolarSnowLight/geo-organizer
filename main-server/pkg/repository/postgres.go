package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable        = "users"
	usersDataTable    = "users_data"
	usersRolesTable   = "users_roles"
	rolesTable        = "roles"
	rolesModulesTable = "roles_modules"
	activationsTable  = "activations"
	tokensTable       = "tokens"
	todoListsTable    = "todo_lists"
	usersListsTable   = "users_lists"
	todoItemsTable    = "todo_items"
	listsItemsTable   = "lists_items"
)

// Структура содержащая параметры подключения к базе данных
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Открытие подключения к базе данных
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CheckRowExists(db *sqlx.DB, table, column, value string) bool {
	query := fmt.Sprintf(`SELECT * FROM %s tl WHERE tl.%s = $1 limit 1`, table, column)
	row := db.QueryRow(query, value)

	var tmp interface{}

	err := row.Scan(&tmp)
	if err != sql.ErrNoRows {
		return true
	}

	return false
}