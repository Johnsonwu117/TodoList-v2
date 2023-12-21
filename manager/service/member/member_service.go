package member

import (
	"log"
	model "todoList/entity/db/member"
	entity "todoList/entity/member"
)

func Create(input *model.Create) (output *model.Base, err error) {
	err = entity.Create(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}
