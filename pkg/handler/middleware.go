package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CtxValue int

const (
	ctxUserID CtxValue = iota
)

func loggingMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"uri":    c.Request.RequestURI,
		}).Info()

		c.Next()
	})
}

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		token, err := getTokenFromRequest(c)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		userId, err := h.usersService.ParseToken(c, token)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(c, ctxUserID, userId)
		c.Request.WithContext(ctx)

		c.Next()
	})
}

func getTokenFromRequest(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}
