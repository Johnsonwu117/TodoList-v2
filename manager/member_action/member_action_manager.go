package member_action

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
	model "todoList/entity/db/list"
	service "todoList/manager/service/member_action"
)

var rdb *redis.Client

func CrateList(input model.Create) any {

	output, err := service.Create(&input)
	if err != nil {
		log.Println("manger is wrong")
	}
	return output
}

func GetList(input model.Search, rdb *redis.Client) any {
	//pkg.SendMail()
	//pkg.SendMailAsync(make(chan bool),input)
	outputTest := &model.Base{}
	redisKey := strconv.Itoa(*input.ID)
	outputRedis, err := rdb.Get(context.Background(), redisKey).Result()
	if err == nil {
		// 將字符串轉換為 JSON 物件
		var outputJSON interface{}
		if err := json.Unmarshal([]byte(outputRedis), &outputJSON); err != nil {
			log.Println(err)
		}

		// 將 JSON 物件轉換為你的 Base 結構
		if jsonData, err := json.Marshal(outputJSON); err == nil {
			if err := json.Unmarshal(jsonData, &outputTest); err != nil {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}

		return outputTest
	}

	output, err := service.GetByID(&input)
	if err != nil {
		log.Println("manger is wrong")
	}
	idStr := strconv.Itoa(output.ID)
	//// 將 output 轉換成 JSON 格式
	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Fatalf("無法將 output 序列化為 JSON：%v", err)
	}
	//// 將 output 存入 Redis，以 id 為 key，設定存活時間為 30 秒
	err = rdb.Set(context.Background(), idStr, outputJSON, 30*time.Second).Err()
	if err != nil {
		log.Fatalf("無法將資料存入 Redis：%v", err)
	}

	return output

}

func UpdateList(input model.Update) any {

	err := service.Update(&input)
	if err != nil {
		log.Println("manger", err)
	}
	return err

}

func DeleteList(input model.Search) any {

	err := service.Delete(&input)
	if err != nil {
		log.Println("manger is wrong")
	}
	return err

}
