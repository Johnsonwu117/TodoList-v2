package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
	"strings"
	memberDB "todoList/entity/db/member"
	dbs "todoList/middleware/connect"
)

var jwtKey = []byte(os.Getenv("JWT_KEY")) // 這是你的密鑰

// 中間函數用於檢查權限
func AuthMiddleware(enforcer *casbin.Enforcer, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			return
		}

		// 解析JWT
		token, err := parseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		// 檢查用戶權限
		currentUser := token.Claims.(jwt.MapClaims)["sub"].(string)
		requestedResource := "/v1/todoList/"

		var members memberDB.Table
		if result := dbs.DB.Where("username = ?", currentUser).First(&members); result.Error != nil {
			log.Println("Error fetching user:", result.Error)
		} else if result.RowsAffected == 0 {
			log.Println("No user found with username:", currentUser)
		} else {
			log.Println("identify", members.Identify)
		}

		if members.Identify == "vip" {
			// 如果是 VIP 用户，允許 POST 請求
			ok, err := enforcer.Enforce(currentUser, requestedResource, "GET")
			if err != nil {
				log.Println("Error enforcing policy:", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
				return
			}
			if ok {
				c.Set("rdb", rdb)
				c.Next()
				return
			}
		} else if members.Identify == "not" {
			if ok, err := enforcer.Enforce(currentUser, requestedResource, "POST"); ok {
				if err != nil {
					log.Println("Error enforcing policy:", err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
					return
				}
				c.Set("rdb", rdb)
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			}
		} else if members.Identify == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		}
	}
}

func parseToken(tokenString string) (*jwt.Token, error) {
	// 將 token 字串以空格拆分為兩部分
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid token format")
	}
	// 選擇第二部分進行解析
	jwtToken := parts[1]

	// 解析JWT
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err // 輸出解析錯誤
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token") // 提示 token 無效
	}

	return token, nil
}
