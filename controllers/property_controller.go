package controllers

import (
	"strconv"
	"time"

	"natsas-place/middleware"
	"natsas-place/models"

	property_service "natsas-place/services/property"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GetAllProperties busca todas as propriedades
func GetAllProperties(c *fiber.Ctx) error {

	var searchInput models.PropertySearchInput

	if err := c.QueryParser(&searchInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid search parameters",
		})
	}

	properties, err := property_service.GetAllPropertiesService(searchInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch properties",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   properties,
	})
}

// GetPropertyById busca uma propriedade por ID
func GetPropertyById(c *fiber.Ctx) error {
	id := c.Params("id")

	property, err := property_service.GetPropertyById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Property not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   property,
	})
}

// CreateProperty cria uma nova propriedade
func CreateProperty(c *fiber.Ctx) error {
	var input models.PropertyCreateInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Validação dos dados
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validation failed",
			"details": err.Error(),
		})
	}

	// Obter usuário atual
	currentUserID := middleware.GetCurrentUser(c)

	// Criar propriedade
	property := models.Property{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Location:    input.Location,
		PriceType:   input.PriceType,
		Bedrooms:    input.Bedrooms,
		Bathrooms:   input.Bathrooms,
		Area:        input.Area,
		Type:        input.Type,
		Amenities:   input.Amenities,
		OwnerID:     currentUserID,
		Status:      "available",
		Available:   true,
		Featured:    false,
	}

	property, err := property_service.CreatePropertyService(property)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create property",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Property created successfully",
		"data":    property,
	})
}

// UpdateProperty atualiza uma propriedade
func UpdateProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var input models.PropertyUpdateInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Validação dos dados
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validation failed",
			"details": err.Error(),
		})
	}

	// Buscar propriedade
	property, err := property_service.GetPropertyById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Property not found",
		})
	}

	// Verificar se o usuário é o proprietário
	currentUserID := middleware.GetCurrentUser(c)
	if property.OwnerID != currentUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "You can only update your own properties",
		})
	}

	// Atualizar campos
	if input.Title != "" {
		property.Title = input.Title
	}
	if input.Description != "" {
		property.Description = input.Description
	}
	if input.Price != nil {
		property.Price = *input.Price
	}
	if input.PriceType != "" {
		property.PriceType = input.PriceType
	}
	if input.Bedrooms != nil {
		property.Bedrooms = *input.Bedrooms
	}
	if input.Bathrooms != nil {
		property.Bathrooms = *input.Bathrooms
	}
	if input.Area != nil {
		property.Area = *input.Area
	}
	if input.Type != "" {
		property.Type = input.Type
	}
	if len(input.Amenities) > 0 {
		property.Amenities = input.Amenities
	}
	if input.Status != "" {
		property.Status = input.Status
	}
	if input.Available != nil {
		property.Available = *input.Available
	}
	if input.Featured != nil {
		property.Featured = *input.Featured
	}
	if input.Location != nil {
		property.Location = *input.Location
	}

	// Salvar alterações
	property, err = property_service.UpdateProperty(property)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update property",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Property updated successfully",
		"data":    property,
	})
}

// DeleteProperty deleta uma propriedade
func DeleteProperty(c *fiber.Ctx) error {
	id := c.Params("id")

	// Buscar propriedade
	property, err := property_service.GetPropertyById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Property not found",
		})
	}

	// Verificar se o usuário é o proprietário
	currentUserID := middleware.GetCurrentUser(c)
	if property.OwnerID != currentUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "You can only delete your own properties",
		})
	}

	// Deletar propriedade
	err = property_service.DeleteProperty(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete property",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Property deleted successfully",
	})
}

// GetMyProperties busca propriedades do usuário atual
func GetMyProperties(c *fiber.Ctx) error {
	currentUserID := middleware.GetCurrentUser(c)

	properties, err := property_service.GetPropertiesByOwner(currentUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch your properties",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   properties,
	})
}

// AddPropertyImage adiciona uma imagem à propriedade
func AddPropertyImage(c *fiber.Ctx) error {
	propertyID := c.Params("id")
	var input models.PropertyImageInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Validação dos dados
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validation failed",
			"details": err.Error(),
		})
	}

	// Verificar se a propriedade existe e pertence ao usuário
	property, err := property_service.GetPropertyById(propertyID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Property not found",
		})
	}

	currentUserID := middleware.GetCurrentUser(c)
	if property.OwnerID != currentUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "You can only add images to your own properties",
		})
	}

	// Converter ID para uint
	propertyIDUint, err := strconv.ParseUint(propertyID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid property ID",
		})
	}

	image, err := property_service.AddPropertyImage(uint(propertyIDUint), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to add property image",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Property image added successfully",
		"data":    image,
	})
}

// RemovePropertyImage remove uma imagem da propriedade
func RemovePropertyImage(c *fiber.Ctx) error {
	imageID := c.Params("imageId")

	// Converter ID para uint
	imageIDUint, err := strconv.ParseUint(imageID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid image ID",
		})
	}

	err = property_service.RemovePropertyImage(uint(imageIDUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to remove property image",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Property image removed successfully",
	})
}

// BookProperty reserva uma propriedade
func BookProperty(c *fiber.Ctx) error {
	propertyID := c.Params("id")
	var input models.BookedPeriodInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Validação dos dados
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validation failed",
			"details": err.Error(),
		})
	}

	// Verificar se as datas são válidas
	if input.StartDate.Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Start date cannot be in the past",
		})
	}

	if input.EndDate.Before(input.StartDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "End date must be after start date",
		})
	}

	// Converter ID para uint
	propertyIDUint, err := strconv.ParseUint(propertyID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid property ID",
		})
	}

	// Definir o guest como o usuário atual
	currentUserID := middleware.GetCurrentUser(c)
	input.GuestID = currentUserID

	booking, err := property_service.BookProperty(uint(propertyIDUint), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Property booked successfully",
		"data":    booking,
	})
}

// GetPropertyBookings busca as reservas de uma propriedade
func GetPropertyBookings(c *fiber.Ctx) error {
	propertyID := c.Params("id")

	// Converter ID para uint
	propertyIDUint, err := strconv.ParseUint(propertyID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid property ID",
		})
	}

	bookings, err := property_service.GetPropertyBookings(uint(propertyIDUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch property bookings",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   bookings,
	})
}

// UpdateBookingStatus atualiza o status de uma reserva
func UpdateBookingStatus(c *fiber.Ctx) error {
	bookingID := c.Params("bookingId")
	status := c.Params("status")

	// Validar status
	validStatuses := []string{"pending", "confirmed", "cancelled", "completed"}
	isValid := false
	for _, validStatus := range validStatuses {
		if status == validStatus {
			isValid = true
			break
		}
	}

	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid status. Must be one of: pending, confirmed, cancelled, completed",
		})
	}

	// Converter ID para uint
	bookingIDUint, err := strconv.ParseUint(bookingID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid booking ID",
		})
	}

	booking, err := property_service.UpdateBookingStatus(uint(bookingIDUint), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update booking status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Booking status updated successfully",
		"data":    booking,
	})
}

// GetAvailableProperties busca propriedades disponíveis para um período
func GetAvailableProperties(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Start date and end date are required",
		})
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid start date format. Use YYYY-MM-DD",
		})
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid end date format. Use YYYY-MM-DD",
		})
	}

	if endDate.Before(startDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "End date must be after start date",
		})
	}

	properties, err := property_service.GetAvailableProperties(startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch available properties",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   properties,
	})
}

// RateProperty avalia uma propriedade
func RateProperty(c *fiber.Ctx) error {
	propertyID := c.Params("id")
	ratingStr := c.Params("rating")

	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid rating value",
		})
	}

	if rating < 0 || rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Rating must be between 0 and 5",
		})
	}

	// Converter ID para uint
	propertyIDUint, err := strconv.ParseUint(propertyID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid property ID",
		})
	}

	err = property_service.UpdatePropertyRating(uint(propertyIDUint), rating)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update property rating",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Property rated successfully",
	})
}
