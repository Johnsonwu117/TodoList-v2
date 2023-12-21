package member

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model "todoList/entity/db/member"
	"todoList/manager/member"
)

// CreateMember
// @Summary Create a new Member
// @Description 新人員註冊
// @Tags Member
// @Accept json
// @param * body model.Create true "新增人員結構檔"
// @Success 200 {string} string "ok"
// @Router /v1/user/register [post]
func CreateMember(c *gin.Context) {
	user := model.Create{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	output := member.CrateMember(user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": output,
	})
}
