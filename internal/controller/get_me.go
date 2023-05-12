package controller

import (
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	a := c.Request.Header.Get("Authorization")
	jwtString := a[7:]
	claims, err := jwt_helper.ParseJWT(jwtString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if userId, ok := claims["user_id"].(int32); ok {
		q := database.NewQuery()
		user, err := q.FindUser(c, userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"resource": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id is invalid"})
	}
}
