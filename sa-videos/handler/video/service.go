package video

import (
	"github.com/Yangusik/sa_videos/models"
	"github.com/gofiber/fiber/v2"
)

const (
	Uploaded   = "Uploaded"
	Processed  = "Processed"
	Processing = "Processing"
	Rejected   = "Rejected"
	Terminated = "Terminated"
	Failed     = "Failed"
	Deleted    = "Deleted"
)

const (
	VideoType  = "Video"
	StreamType = "Stream"
)

func saveFileTemp(ctx *fiber.Ctx, key string) (details models.FileDetails, err error) {
	fileDetails := models.FileDetails{}

	file, err := ctx.FormFile(key)
	if err != nil {
		return fileDetails, err
	}

	path := "storage/temp/" + file.Filename

	fileDetails.Name = file.Filename
	fileDetails.Size = uint(file.Size)
	fileDetails.Type = VideoType
	fileDetails.Path = path
	//fileDetails.VideoStreams
	fileDetails.DurationMs = 100

	if err := ctx.SaveFile(file, path); err != nil {
		return fileDetails, err
	}

	return fileDetails, nil
}
