package member_action

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	middle "todoList/middleware/auth"
	"todoList/presenter/member_action"
)

func GetRoute(r *gin.RouterGroup, enforcer *casbin.Enforcer, rdb *redis.Client) {

	user := r.Group("/todoList")
	user.GET("/", middle.AuthMiddleware(enforcer, rdb), member_action.GetList)
	user.POST("/", middle.AuthMiddleware(enforcer, rdb), member_action.CreateList)
	user.PUT("/:id", middle.AuthMiddleware(enforcer, rdb), member_action.UpdateList)
	user.DELETE("/:id", middle.AuthMiddleware(enforcer, rdb), member_action.DeleteList)
}
