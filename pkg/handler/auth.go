package handler

import (
	"ImGoDeveloper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input ImGoDeveloper.User

	if err := c.BindJSON(&input); err != nil {
		newErrrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {

}
