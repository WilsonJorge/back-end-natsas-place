package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street       string  `json:"street" validate:"required,min=5,max=200"`
	Number       string  `json:"number" validate:"required,min=1,max=20"`
	Neighborhood string  `json:"neighborhood" validate:"required,min=2,max=100"`
	City         string  `json:"city" validate:"required,min=2,max=100"`
	State        string  `json:"state" validate:"required,min=2,max=100"`
	Country      string  `json:"country" validate:"required,min=2,max=100"`
	ZipCode      string  `json:"zip_code" validate:"required,min=5,max=20"`
	Latitude     float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude    float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Instructions string  `json:"instructions" gorm:"type:text"`
}

type AddressCreateInput struct {
	Street       string  `json:"street" validate:"required,min=5,max=200"`
	Number       string  `json:"number" validate:"required,min=1,max=20"`
	Neighborhood string  `json:"neighborhood" validate:"required,min=2,max=100"`
	City         string  `json:"city" validate:"required,min=2,max=100"`
	State        string  `json:"state" validate:"required,min=2,max=100"`
	Country      string  `json:"country" validate:"required,min=2,max=100"`
	ZipCode      string  `json:"zip_code" validate:"required,min=5,max=20"`
	Latitude     float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude    float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Instructions string  `json:"instructions"`
}

type AddressUpdateInput struct {
	Street       string  `json:"street" validate:"min=5,max=200"`
	Number       string  `json:"number" validate:"min=1,max=20"`
	Neighborhood string  `json:"neighborhood" validate:"min=2,max=100"`
	City         string  `json:"city" validate:"min=2,max=100"`
	State        string  `json:"state" validate:"min=2,max=100"`
	Country      string  `json:"country" validate:"min=2,max=100"`
	ZipCode      string  `json:"zip_code" validate:"min=5,max=20"`
	Latitude     float64 `json:"latitude" validate:"gte=-90,lte=90"`
	Longitude    float64 `json:"longitude" validate:"gte=-180,lte=180"`
	Instructions string  `json:"instructions"`
}

type AddressResponse struct {
	ID           uint    `json:"id"`
	Street       string  `json:"street"`
	Number       string  `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	ZipCode      string  `json:"zip_code"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Instructions string  `json:"instructions"`
	FullAddress  string  `json:"full_address"`
}

type GeocodeInput struct {
	Address string `json:"address" validate:"required,min=10,max=500"`
}

type GeocodeResponse struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Street       string  `json:"street"`
	Number       string  `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	ZipCode      string  `json:"zip_code"`
	FullAddress  string  `json:"full_address"`
}

type LocationSearchInput struct {
	Latitude  float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	RadiusKm  float64 `json:"radius_km" validate:"required,gt=0,lte=100"`
	Query     string  `json:"query"`
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
}
