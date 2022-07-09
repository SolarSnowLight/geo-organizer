package handler

import (
	"errors"
	middlewareConstants "main-server/pkg/constants/middleware"
	"main-server/pkg/service/google_oauth2"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(middlewareConstants.AUTHORIZATION_HEADER)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Пустой заголовок авторизации!")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Не корректный авторизационный заголовок!")
		return
	}

	data, err := h.services.Token.ParseToken(headerParts[1], viper.GetString("token.signing_key_access"))

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	switch data.AuthType.Value {
	case "GOOGLE":
		if result, err := google_oauth2.VerifyAccessToken(*data.TokenApi); err != nil || result != true {
			newErrorResponse(c, http.StatusUnauthorized, "Не действительный токен доступа")
			return
		}
		break

	case "LOCAL":
		break
	}

	// Добавление к контексту дополнительных данных о пользователе
	c.Set(middlewareConstants.USER_CTX, data.UsersId)
	c.Set(middlewareConstants.ROLES_CTX, data.RolesId)
	c.Set(middlewareConstants.AUTH_TYPE_VALUE_CTX, data.AuthType.Value)
	c.Set(middlewareConstants.TOKEN_API_CTX, data.TokenApi)
	c.Set(middlewareConstants.ACCESS_TOKEN_CTX, headerParts[1])
}

func (h *Handler) userIdentityLogout(c *gin.Context) {
	header := c.GetHeader(middlewareConstants.AUTHORIZATION_HEADER)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Пустой заголовок авторизации!")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Не корректный авторизационный заголовок!")
		return
	}

	data, err := h.services.Token.ParseTokenWithoutValid(headerParts[1], viper.GetString("token.signing_key_access"))

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Добавление к контексту дополнительных данных о пользователе
	c.Set(middlewareConstants.USER_CTX, data.UsersId)
	c.Set(middlewareConstants.ROLES_CTX, data.RolesId)
	c.Set(middlewareConstants.AUTH_TYPE_VALUE_CTX, data.AuthType.Value)
	c.Set(middlewareConstants.TOKEN_API_CTX, data.TokenApi)
	c.Set(middlewareConstants.ACCESS_TOKEN_CTX, headerParts[1])
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(middlewareConstants.USER_CTX)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
