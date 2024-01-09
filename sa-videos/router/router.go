package router

import (
	"github.com/Yangusik/sa_videos/handler"
	"github.com/Yangusik/sa_videos/handler/video"
	"github.com/Yangusik/sa_videos/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/videos", logger.New())
	apiAdmin := api.Group("/admin", logger.New(), middleware.AuthAdmin())

	api.Get("/", handler.Hello)
	api.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{"message": "Ok"})
	})

	api.Get("/view", video.ListVideos)
	api.Get("/get", video.GetVideo)
	api.Post("/upload", video.UploadVideo)
	api.Delete("/deleteALl", video.DeleteAllVideos)

	apiAdmin.Get("/", handler.Hello)
}
