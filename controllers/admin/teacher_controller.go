package controller_admin

import "github.com/gin-gonic/gin"

func CreateTeacher(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "hello",
	})
}
