package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID                uint              `json:"id" gorm:"primarykey"`
	Title             string            `json:"title" gorm:"type:text;not null;"`
	Description       string            `json:"description"`
	UserId            uint              `json:"user_id"`
	Status            string            `json:"status"`
	FileDetails       FileDetails       `json:"fileDetails" gorm:"serializer:json"`
	ProcessingDetails ProcessingDetails `json:"processingDetails" gorm:"serializer:json"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt    `json:"deletedAt" gorm:"index"`
}

type FileDetails struct {
	Name         string       `json:"name"`
	Path         string       `json:"path"`
	Size         uint         `json:"size"`
	Type         string       `json:"type"`
	VideoStreams VideoStreams `json:"videoStreams" gorm:"serializer:json"`
	DurationMs   uint         `json:"durationMs"`
}

type VideoStreams struct {
	WidthPixels  uint    `json:"widthPixels"`
	HeightPixels uint    `json:"heightPixels"`
	FrameRateFps float32 `json:"frameRateFps"`
	AspectRatio  float32 `json:"aspectRatio"`
	Codec        string  `json:"codec"`
	BitrateBps   float32 `json:"bitrateBps"`
	Rotation     string  `json:"rotation"`
}

type ProcessingDetails struct {
	Status   string   `json:"status"`
	Progress Progress `json:"progress"  gorm:"serializer:json"`
}

type Progress struct {
	PartsTotal     float32 `json:"partsTotal"`
	PartsProcessed float32 `json:"partsProcessed"`
	TimeLeftMs     float32 `json:"timeLeftMs"`
}
