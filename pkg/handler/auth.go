package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SenselessA/CRUD_books"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {

	var inp CRUD_books.SignUpInput
	if err := c.BindJSON(&inp); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.User.SignUp(c, inp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) SignIn(c *gin.Context) {
	var inp CRUD_books.SignInInput

	if err := c.BindJSON(&inp); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := h.User.SignIn(c, inp)
	if err != nil {
		if errors.Is(err, CRUD_books.ErrUserNotFound) {
			handleNotFoundError(c, err)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func handleNotFoundError(c *gin.Context, err error) {
	response, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})

	c.JSON(http.StatusBadRequest, response)
}
