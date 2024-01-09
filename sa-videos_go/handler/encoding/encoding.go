package encoding

import (
	"bytes"
	"github.com/Yangusik/sa_videos/models"
	"github.com/streadway/amqp"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"path/filepath"
	"time"
)

func EncodeVideo(msg amqp.Delivery, video *models.Video) {
	log.Println(time.Now().Format("01-02-2006 15:04:05"), "::", string(bytes.Clone(msg.Body)))

	inputPath, _ := filepath.Abs(video.FileDetails.Path)
	outputPath, _ := filepath.Abs("storage/encoding/" + video.FileDetails.Name)

	if err := ffmpeg.Input(inputPath).
		Output(outputPath, ffmpeg.KwArgs{"c:v": "libvpx-vp9"}).
		OverWriteOutput().
		ErrorToStdOut().
		Run(); err != nil {
		log.Fatal(err)
	}

	//msg.Ack(false)
}
