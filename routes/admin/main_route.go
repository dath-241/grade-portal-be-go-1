package routes_admin

<<<<<<< HEAD
import (
	"LearnGo/config"
	middlewares_admin "LearnGo/middlewares/admin"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.Engine) {
	prefixAdmin := config.PrefixAdmin()
	AuthRoute(r.Group(prefixAdmin))
	// middleware đảm bảo rằng đã đăng nhập trước khi vào web
	protectedGroup := r.Group(prefixAdmin)
	protectedGroup.Use(middlewares_admin.RequireAuth)
	// tạo các group để chạy các api sau khi đã đăng nhập thành công
	TeacherRoute(protectedGroup.Group("/teacher"))
=======
import "github.com/gin-gonic/gin"

func MainRoute(r *gin.Engine) {
	AuthRoute(r.Group("/admin/api"))
>>>>>>> d02509ed7fcdab80770194afdfcb89e5f7eae356
}
