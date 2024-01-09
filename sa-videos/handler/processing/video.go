package processing

import (
	"encoding/json"
	"github.com/Yangusik/sa_videos/handler/encoding"
	videoPackage "github.com/Yangusik/sa_videos/handler/video"
	"github.com/Yangusik/sa_videos/models"
	"github.com/streadway/amqp"
	"log"
)

func Processing(msg amqp.Delivery) {
	video := models.Video{}
	if err := json.Unmarshal(msg.Body, &video); err != nil {
		log.Fatal(err)
	}

	switch video.Status {
	case videoPackage.Uploaded:
		encoding.EncodeVideo(msg, &video)
	default:
		log.Fatalf("Not found action for status", video.Status)
	}
}
