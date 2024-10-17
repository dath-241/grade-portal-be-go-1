package controller_admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	fmt.Println(c.Get("ID"))
	c.JSON(200, gin.H{
		"code": "hello",
	})
}
