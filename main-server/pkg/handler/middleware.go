package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	authorizationHeader = "Authorization"
	usersCtx            = "users_id"
	rolesCtx            = "roles_id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Пустой заголовок авторизации!")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Не корректный авторизационный заголовок!")
		return
	}

	// Парсинг токена доступа
	data, err := h.services.Authorization.ParseToken(headerParts[1], viper.GetString("token.signing_key_access"))

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(usersCtx, data.UsersId)
	c.Set(rolesCtx, data.RolesId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(usersCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
