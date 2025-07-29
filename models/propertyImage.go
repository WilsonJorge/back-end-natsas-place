package models

import (
	"time"

	"gorm.io/gorm"
)

type PropertyImage struct {
	gorm.Model
	PropertyID   uint      `json:"property_id" gorm:"not null"`
	ImageURL     string    `json:"image_url" gorm:"not null" validate:"required,url"`
	ThumbnailURL string    `json:"thumbnail_url" validate:"omitempty,url"`
	Title        string    `json:"title" validate:"max=100"`
	Description  string    `json:"description" validate:"max=500"`
	AltText      string    `json:"alt_text" validate:"max=200"`
	IsMain       bool      `json:"is_main" gorm:"default:false"`
	Order        int       `json:"order" gorm:"default:0" validate:"min=0,max=100"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	FileSize     int64     `json:"file_size"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Format       string    `json:"format"`
	UploadedAt   time.Time `json:"uploaded_at" gorm:"autoCreateTime"`

	// Relationships
	Property Property `json:"property" gorm:"foreignKey:PropertyID"`
}

type PropertyImageCreateInput struct {
	PropertyID   uint   `json:"property_id" validate:"required"`
	ImageURL     string `json:"image_url" validate:"required,url"`
	ThumbnailURL string `json:"thumbnail_url" validate:"omitempty,url"`
	Title        string `json:"title" validate:"max=100"`
	Description  string `json:"description" validate:"max=500"`
	AltText      string `json:"alt_text" validate:"max=200"`
	IsMain       bool   `json:"is_main"`
	Order        int    `json:"order" validate:"min=0,max=100"`
	IsActive     bool   `json:"is_active"`
}

type PropertyImageUpdateInput struct {
	ImageURL     string `json:"image_url" validate:"omitempty,url"`
	ThumbnailURL string `json:"thumbnail_url" validate:"omitempty,url"`
	Title        string `json:"title" validate:"max=100"`
	Description  string `json:"description" validate:"max=500"`
	AltText      string `json:"alt_text" validate:"max=200"`
	IsMain       bool   `json:"is_main"`
	Order        int    `json:"order" validate:"min=0,max=100"`
	IsActive     bool   `json:"is_active"`
}

type PropertyImageResponse struct {
	ID           uint      `json:"id"`
	PropertyID   uint      `json:"property_id"`
	ImageURL     string    `json:"image_url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	AltText      string    `json:"alt_text"`
	IsMain       bool      `json:"is_main"`
	Order        int       `json:"order"`
	IsActive     bool      `json:"is_active"`
	FileSize     int64     `json:"file_size"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Format       string    `json:"format"`
	UploadedAt   time.Time `json:"uploaded_at"`
}

type ImageUploadInput struct {
	PropertyID  uint   `json:"property_id" validate:"required"`
	Title       string `json:"title" validate:"max=100"`
	Description string `json:"description" validate:"max=500"`
	AltText     string `json:"alt_text" validate:"max=200"`
	IsMain      bool   `json:"is_main"`
	Order       int    `json:"order" validate:"min=0,max=100"`
}

type ImageUploadResponse struct {
	ID           uint   `json:"id"`
	ImageURL     string `json:"image_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	FileSize     int64  `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Format       string `json:"format"`
	Message      string `json:"message"`
}

type BulkImageUpdateInput struct {
	PropertyID uint                       `json:"property_id" validate:"required"`
	Images     []PropertyImageUpdateInput `json:"images" validate:"required,min=1"`
}

type ImageReorderInput struct {
	PropertyID uint              `json:"property_id" validate:"required"`
	Images     []ImageOrderInput `json:"images" validate:"required,min=1"`
}

type ImageOrderInput struct {
	ID    uint `json:"id" validate:"required"`
	Order int  `json:"order" validate:"min=0,max=100"`
}
