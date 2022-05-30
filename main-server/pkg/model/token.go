package model

type TokenModel struct {
	Id           int    `json:"id" db:"id"`
	UsersId      int    `json:"users_id" db:"users_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type TokenDataModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenOutputParse struct {
	UsersId int `json:"users_id"`
	RolesId int `json:"roles_id"`
}

type TokenOutputParseString struct {
	UsersId string `json:"users_id"`
	RolesId string `json:"roles_id"`
}

type TokenRefreshModel struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
