package main

import (
	"fmt"
	"github.com/Yangusik/sa_videos/database"
	"github.com/Yangusik/sa_videos/handler/processing"
	"github.com/Yangusik/sa_videos/queue"
	"github.com/Yangusik/sa_videos/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading  .env file:", err)
	}

	database.ConnectDB()
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})
	router.SetupRoutes(app)
	queue.Init(fmt.Sprintf("amqp://%s:%s@rabbitmq", os.Getenv("RABBITMQ_DEFAULT_USER"), os.Getenv("RABBITMQ_DEFAULT_PASS")))
	go queueWorker()
	app.Listen(":3000")
}

func queueWorker() {
	msgs, close, err := queue.Subscribe("video")
	if err != nil {
		panic(err)
	}
	defer close()

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			processing.Processing(msg)
		}
	}()
	<-forever
}
