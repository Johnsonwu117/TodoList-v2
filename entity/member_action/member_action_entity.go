package member_action

import (
	"gorm.io/gorm/clause"
	"log"
	model "todoList/entity/db/list"
	db "todoList/middleware/connect"
)

func Create(input *model.Table) (err error) {
	err = db.DB.Model(&model.Table{}).Omit(clause.Associations).Create(&input).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetByID(input *model.Search) (output *model.Table, err error) {
	dbs := db.DB.Model(&model.Table{})

	if input.ID != nil {
		dbs.Where("id = ?", input.ID)
	}

	if input.Completed != nil {
		dbs.Where("completed = ?", input.Completed)

	}

	if input.Title != nil {
		dbs.Where("title like ?", "%"+*input.Title+"%")
	}

	if input.Description != nil {
		dbs.Where("description LIKE ?", "%"+*input.Description+"%")
	}

	if input.Priority != nil {
		dbs.Where("priority = ?", input.Priority)
	}

	err = dbs.First(&output).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func Update(input *model.Update) (err error) {
	dbs := db.DB.Model(&model.Table{})
	data := map[string]any{}

	if input.ID != nil {
		dbs.Where("id = ?", input.ID)
	}

	if input.Completed != nil {
		data["completed"] = input.Completed
	}

	if input.Priority != nil {
		data["priority"] = input.Priority
	}

	if input.Title != nil {
		data["title"] = input.Title
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.DueDate != nil {
		data["duedate"] = input.DueDate
	}

	err = dbs.Select("*").Updates(data).Error
	if err != nil {
		log.Println("entity", err)
		return err
	}

	return nil
}

func Delete(input *model.Search) (err error) {
	dbs := db.DB.Model(&model.Table{})

	if input.ID != nil {
		dbs.Where("id = ?", input.ID)
	}

	if input.Completed != nil {
		dbs.Where("completed = ?", input.Completed)

	}

	if input.Title != nil {
		dbs.Where("title like %?%", input.Title)

	}

	if input.Description != nil {
		dbs.Where("description like %?%", input.Description)

	}

	err = dbs.Delete(&model.Table{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
