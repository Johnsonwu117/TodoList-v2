package member_action

import (
	"encoding/json"
	"log"
	"time"
	model "todoList/entity/db/list"
	entity "todoList/entity/member_action"
)

func GetByID(input *model.Search) (output *model.Base, err error) {
	field, err := entity.GetByID(input)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func Create(input *model.Create) (output *model.Base, err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output.CreatedAt = time.Now()
	marshal, err = json.Marshal(output)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	table := &model.Table{}
	err = json.Unmarshal(marshal, &table)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = entity.Create(table)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func Update(input *model.Update) (err error) {
	err = entity.Update(input)
	if err != nil {
		log.Println("service", err)
		return err
	}

	return nil
}

func Delete(input *model.Search) (err error) {
	err = entity.Delete(input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
