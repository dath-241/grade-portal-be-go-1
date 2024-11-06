package controller_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"author":   "Thai Ly",
		"facebook": "https://www.facebook.com/thai.ly.716970",
		"contact":  "lyvinhthai321@gmail.com",
	})
}
