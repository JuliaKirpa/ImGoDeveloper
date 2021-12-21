package handler

import (
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
