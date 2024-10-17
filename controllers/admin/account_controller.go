package controller_admin

import (
	"github.com/gin-gonic/gin"
)

func AccountCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "da vao trang account create",
	})
}
