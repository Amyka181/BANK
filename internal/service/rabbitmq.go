package service

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func SendToRabbit(user entity.UpdateUser) {
	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Println("Не удалось подключиться к RabbitMQ: %v", err)
	} else {
		log.Println("Успешно подключено к RabbitMQ")
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Println("Не удалось открыть канал: %v", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"my_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Не удалось создать очередь: %v", err)
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println("Ошибка при кодировании в Json")
	}
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Не удалось отправить сообщение: %v", err)
	}

	log.Println("Сообщение отправлено в очередь!")

}
