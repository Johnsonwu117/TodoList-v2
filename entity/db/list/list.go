package list

import (
	"time"
)

type Table struct {
	ID int `json:"id"`
	//標題
	Title string `json:"title"`
	//事項說明
	Description string `json:"description"`
	//截止日期
	DueDate string `json:"duedate" gorm:"column:duedate"`
	//事項優先級
	Priority int `json:"priority"`
	//是否完成
	Completed bool `json:"completed"`

	// 創建時間
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdat"`
	// 更新時間
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedat"`
	// 刪除時間
	DeletedAt time.Time `json:"deletedAt,omitempty" gorm:"column:deletedat"`
}

type Base struct {
	ID int `json:"id"`
	//標題
	Title string `json:"title"`
	//事項說明
	Description string `json:"description"`
	//截止日期
	DueDate string `json:"duedate"`
	//事項優先級
	Priority int `json:"priority"`
	//是否完成
	Completed bool `json:"completed"`

	// 創建時間
	CreatedAt time.Time `json:"createdAt"`
	// 更新時間
	UpdatedAt time.Time `json:"updatedAt"`
	// 刪除時間
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type Search struct {
	//事項id
	ID *int `json:"id" form:"id"`
	//標題
	Title *string `json:"title" form:"title"`
	//事項說明
	Description *string `json:"description" form:"description"`
	//截止日期
	DueDate *string `json:"duedate" form:"duedate"`
	//事項優先級
	Priority *int `json:"priority" form:"priority"`
	//是否完成
	Completed *bool `json:"completed" form:"completed"`
}

type Create struct {
	//標題
	Title string `json:"title"`
	//事項說明
	Description string `json:"description"`
	//截止日期
	DueDate string `json:"duedate"`
	//事項優先級
	Priority int `json:"priority"`
	//是否完成
	Completed bool `json:"completed"`
}

type Update struct {
	//事項id
	ID *int `json:"id"`
	//標題
	Title *string `json:"title"`
	//事項說明
	Description *string `json:"description"`
	//截止日期
	DueDate *string `json:"duedate"`
	//事項優先級
	Priority *int `json:"priority"`
	//是否完成
	Completed *bool `json:"completed"`
}

func (t *Table) TableName() string {
	return "lists"
}
