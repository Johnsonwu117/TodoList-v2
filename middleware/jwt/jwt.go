package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	memberDB "todoList/entity/db/member"
	dbs "todoList/middleware/connect"
)

type CustomClaims struct {
	jwt.StandardClaims
	Subject string `json:"sub"`
}

var jwtKey = []byte(os.Getenv("JWT_KEY")) // 這密鑰

// 登入
func Login(c *gin.Context) {
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// 查db有無資料
	var members memberDB.Table
	result := dbs.DB.Where("username = ?", creds.Username).First(&members)
	if result.RowsAffected == 0 || members.UserPassword != creds.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// 創造 JWT
	expirationTime := time.Now().Add(time.Hour * 12)

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		Subject: creds.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
