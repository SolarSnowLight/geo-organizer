package handler

import (
	"main-server/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	tokenTTL_access  = 1 * time.Hour
	tokenTTL_refresh = 12 * time.Hour
)

// @Summary SignUp
// @Tags auth
// @Description Регистрация пользователя
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.UserRegisterModel true "account info"
// @Success 200 {object} model.UserAuthDataModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input model.UserRegisterModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary SignIn
// @Tags auth
// @Description Авторизация пользователя
// @ID login
// @Accept  json
// @Produce  json
// @Param input body model.UserLoginModel true "credentials"
// @Success 200 {object} model.UserAuthDataModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input model.UserLoginModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Authorization.LoginUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary Refresh
// @Tags auth
// @Description Обновление токена доступа и токена обновления
// @ID refresh
// @Accept  json
// @Produce  json
// @Param input body model.TokenRefreshModel true "credentials"
// @Success 200 {object} model.UserAuthDataModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var input model.TokenRefreshModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Authorization.Refresh(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

type LogoutOutputModel struct {
	IsLogout bool `json:"is_logout"`
}

// @Summary Refresh
// @Tags auth
// @Description Обновление токена доступа и токена обновления
// @ID refresh
// @Accept  json
// @Produce  json
// @Param input body model.TokenDataModel true "credentials"
// @Success 200 {object} LogoutOutputModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh [post]
func (h *Handler) logout(c *gin.Context) {
	var input model.TokenDataModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Authorization.Logout(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, LogoutOutputModel{
		IsLogout: data,
	})
}
