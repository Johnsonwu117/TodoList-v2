package pkg

import (
	"github.com/streadway/amqp"
	"log"
	"net/smtp"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendMail() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // 連接到 RabbitMQ
	failOnError(err, "無法連接到 RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel() // 建立一個通道
	failOnError(err, "無法打開通道")
	defer ch.Close()
	q, err := ch.QueueDeclarePassive(
		"emailQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("隊列不存在，有其他錯誤:", err)
		return
	}

	body := "您已成功註冊,恭喜你加入！" // 註冊成功消息

	err = ch.Publish(
		"",     // 交換器
		q.Name, // 路由鍵
		false,  // 強制性
		false,  // 立即發送
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "無法發送消息")

	log.Println("訊息已成功發送：", body)
}

func SendMailAsync(done chan bool, email string) {

	go func() {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // 連接到 RabbitMQ
		failOnError(err, "無法連接到 RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel() // 建立一個通道
		failOnError(err, "無法打開通道")
		defer ch.Close()

		q, err := ch.QueueDeclarePassive(
			"emailQueue",
			false,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "無法宣告消息隊列")

		msgs, err := ch.Consume(
			q.Name, // 從消息隊列訂閱消息
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "無法訂閱消息")
		for d := range msgs {
			body := string(d.Body) // 獲取郵件正文

			// 以下是發送郵件的代碼
			from := os.Getenv("GMAIL_EMAIL")        // 更換為您的 Gmail 郵箱地址
			password := os.Getenv("GMAIL_PASSWORD") // 更換為您的 Gmail 密碼
			to := email                             // 更換為收件人的電子郵件地址
			subject := "Subject: 註冊成功通知\n"
			msg := []byte(subject + "\r\n" + body)

			err := smtp.SendMail("smtp.gmail.com:587",
				smtp.PlainAuth("", from, password, "smtp.gmail.com"),
				from,
				[]string{to},
				msg,
			)
			failOnError(err, "無法發送郵件")

			log.Printf("已發送郵件: %s\n", body)
		}
		done <- true // 發送結束信號
	}()
}
