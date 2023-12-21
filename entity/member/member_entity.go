package member

import (
	"gorm.io/gorm/clause"
	"log"
	model "todoList/entity/db/member"
	db "todoList/middleware/connect"
)

func Create(input *model.Create) (err error) {
	err = db.DB.Model(&model.Create{}).Omit(clause.Associations).Create(&input).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
