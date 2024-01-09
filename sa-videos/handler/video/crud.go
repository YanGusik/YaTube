package video

import (
	"encoding/json"
	"errors"
	"github.com/Yangusik/sa_videos/database"
	"github.com/Yangusik/sa_videos/helper"
	"github.com/Yangusik/sa_videos/models"
	"github.com/Yangusik/sa_videos/queue"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
)

func ListVideos(ctx *fiber.Ctx) error {
	videos := []models.Video{}
	database.DB.DB.Find(&videos)
	return ctx.Status(200).JSON(videos)
}

func GetVideo(ctx *fiber.Ctx) error {
	video := models.Video{}

	id := ctx.QueryInt("id", 0)
	err := database.DB.DB.Where("id=?", id).First(&video).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(&fiber.Map{"message": "Not Found", "id": id})
	}

	return ctx.Status(200).JSON(video)
}

func UploadVideo(ctx *fiber.Ctx) error {
	fileDetails, err := saveFileTemp(ctx, "upload")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user := helper.GetUserFromContext(ctx)

	//form TODO: validation
	title := ctx.Query("title", "MyFirstVideo")
	description := ctx.Query("description", "Description Sample")

	video := models.Video{
		Title:             title,
		Description:       description,
		UserId:            user.Id,
		Status:            Uploaded,
		FileDetails:       fileDetails,
		ProcessingDetails: models.ProcessingDetails{Status: Uploaded},
	}

	tx := database.DB.DB.Begin()
	if errDb := tx.Create(&video).Error; errDb != nil {
		tx.Rollback()
		errOs := os.Remove(fileDetails.Path)
		if errOs != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors.New(errOs.Error() + errDb.Error()))
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	jsonText, err := json.Marshal(&video)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := queue.Publish("video", jsonText); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	tx.Commit()
	return ctx.Status(200).JSON(&fiber.Map{"message": "Success", "data": video})
}

func DeleteAllVideos(ctx *fiber.Ctx) error {
	user := helper.GetUserFromContext(ctx)
	database.DB.DB.Where("user_id = ?", user.Id).Delete(&models.Video{})
	return ctx.Status(200).JSON(&fiber.Map{"message": "Success"})
}
