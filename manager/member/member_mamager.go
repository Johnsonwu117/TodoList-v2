package member

import (
	"log"
	model "todoList/entity/db/member"
	service "todoList/manager/service/member"
	"todoList/pkg"
)

func CrateMember(input model.Create) any {
	pkg.SendMail()
	output, err := service.Create(&input)
	if err != nil {
		log.Println("manger is wrong")
	}
	//connect.SendMailAsync(make(chan bool), input.UserEmail)
	return output
}
