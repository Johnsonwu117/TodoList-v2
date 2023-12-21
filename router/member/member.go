package member

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	middle "todoList/middleware/jwt"
	members "todoList/presenter/member"
)

func GetRoute(r *gin.RouterGroup, enforcer *casbin.Enforcer, rdb *redis.Client) {

	member := r.Group("/user")
	member.POST("/login", middle.Login)           // 新增登入路由用於獲取 JWT
	member.POST("register", members.CreateMember) //創建新用戶

}
