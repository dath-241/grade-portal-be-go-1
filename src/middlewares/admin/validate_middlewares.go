package middlewares_admin

import (
	controller_admin "LearnGo/controllers/admin"
	"strings"

	"github.com/gin-gonic/gin"
)

//	if !strings.HasSuffix(newUser.Email, "@hcmut.edu.vn") {
//		invalidEmails = append(invalidEmails, newUser)
//		continue
//	}
func ValindateEmail(email string) bool {
	return strings.HasSuffix(email, "@hcmut.edu.vn")
}

func ValindateMS(ms string) bool {
	return ms != ""
}

func ValidateDataAdmin(c *gin.Context) {
	var data controller_admin.InterfaceAdminController
	c.BindJSON(&data)
	if !ValindateEmail(data.Email) || data.Ms == "" {
		c.JSON(400, gin.H{
			"code":    "error",
			"massage": "Dữ liệu không hợp lệ",
			"data":    data,
		})
		c.Abort()
		return
	}
	c.Set("adminData", data)
	c.Next()
}
