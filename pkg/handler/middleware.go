package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AUTHORIZATIONHEADER = "Authorization"
	USERCTX             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AUTHORIZATIONHEADER)
	if header == "" {
		newErrrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(USERCTX, userId)
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(USERCTX)
	if !ok {
		newErrrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrrorResponse(c, http.StatusInternalServerError, "user id has invalid type")
		return 0, errors.New("user id has invalid type")
	}
	return idInt, nil
}
