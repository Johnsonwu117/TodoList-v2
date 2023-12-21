package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"todoList/middleware/connect"
	member "todoList/router/member"
	action "todoList/router/member_action"
)

// @contact.name   todolist demo
// @title todoList demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
func main() {
	connect.ConnectToPostgres()
	rdb := connect.ConnectToRedis()
	router := gin.Default()
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := router.Group("/v1")

	// 創建 Casbin 實例
	m, err := model.NewModelFromFile("auth_model.conf")
	if err != nil {
		log.Fatal(err)
	}

	adapter := fileadapter.NewAdapter("policy.csv")

	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatal(err)
	}

	// 加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatal(err)
	}

	action.GetRoute(v1, enforcer, rdb)
	member.GetRoute(v1, enforcer, rdb)
	router.Run(":8080")
}
