package member_action

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"todoList/entity/db/list"
	"todoList/manager/member_action"
)

// CreateList
// @Summary Create a new To-do list
// @Description 新建待辦事項
// @Tags Member_Action
// @Accept json
// @param * body list.Create true "新增事項結構檔"
// @Success 200 {string} string "ok"
// @Router /v1/todoList/ [post]
func CreateList(c *gin.Context) {
	user := list.Create{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	output := member_action.CrateList(user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": output,
	})
}

// GetList
// @Summary Get a To-do list
// @Description 查詢待辦事項
// @Tags Member_Action
// @Accept json
// @Param id query string false "事項id"
// @Param title query string false "事項"
// @Param description query string false "事項說明"
// @Param title query string false "截止日期"
// @Param completed query string false "是否完成"
// @Param priority query string false "事項優先級"
// @Success 200 {string} string "ok"
// @Router /v1/todoList/ [get]
func GetList(c *gin.Context) {
	user := &list.Search{}
	err := c.ShouldBindQuery(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}
	rdb := c.MustGet("rdb").(*redis.Client)
	output := member_action.GetList(*user, rdb)
	c.JSON(http.StatusOK, gin.H{
		"訊息": output,
	})
}

// UpdateList
// @Summary  Update To-do list
// @Description 更新代辦事項
// @Tags Member_Action
// @Accept json
// @Param id query string false "事項id"
// @param * body list.Update true "修改事項結構檔"
// @Success 200 {string} string "ok"
// @Router /v1/todoList/:id [put]
func UpdateList(c *gin.Context) {
	user := &list.Update{}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	user.ID = &id

	err = c.BindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}
	output := member_action.UpdateList(*user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": output,
	})
}

// DeleteList
// @Summary Get a To-do list
// @Description 刪除代辦事項
// @Tags Member_Action
// @Accept json
// @Param id query string true "事項id"
// @Success 200 {string} string "ok"
// @Router /v1/todoList/:id [delete]
func DeleteList(c *gin.Context) {
	user := &list.Search{}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	user.ID = &id

	err = c.ShouldBindQuery(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	output := member_action.DeleteList(*user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": output,
	})
}
