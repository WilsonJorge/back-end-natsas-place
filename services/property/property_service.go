package property

import (
	"fmt"
	"natsas-place/database"
	"natsas-place/models"
	"time"
)

// CreatePropertyService cria uma nova propriedade
func CreatePropertyService(propertyData models.Property) (models.Property, error) {
	err := database.Instance.Create(&propertyData).Error
	if err != nil {
		return models.Property{}, fmt.Errorf("failed to create property: %w", err)
	}
	return propertyData, nil
}

// GetPropertyById busca uma propriedade por ID
func GetPropertyById(id string) (models.Property, error) {
	var property models.Property
	err := database.Instance.Preload("Owner").Preload("Images").Preload("BookedDates").Where("id = ?", id).First(&property).Error
	if err != nil {
		return models.Property{}, fmt.Errorf("property not found: %w", err)
	}

	return property, nil
}

// GetAllPropertiesService busca todas as propriedades
func GetAllPropertiesService(searchInput models.PropertySearchInput) ([]models.Property, error) {

	var properties []models.Property
	query := database.Instance.Preload("Owner").Preload("Images")

	// Aplicar filtros
	if searchInput.Location != "" {
		query = query.Where("location->>'$.city' LIKE ? OR location->>'$.address' LIKE ?", "%"+searchInput.Location+"%", "%"+searchInput.Location+"%")
	}
	if searchInput.MinPrice != nil {
		query = query.Where("price >= ?", *searchInput.MinPrice)
	}
	if searchInput.MaxPrice != nil {
		query = query.Where("price <= ?", *searchInput.MaxPrice)
	}
	if searchInput.PriceType != "" {
		query = query.Where("price_type = ?", searchInput.PriceType)
	}
	if searchInput.MinBedrooms != nil {
		query = query.Where("bedrooms >= ?", *searchInput.MinBedrooms)
	}
	if searchInput.MaxBedrooms != nil {
		query = query.Where("bedrooms <= ?", *searchInput.MaxBedrooms)
	}
	if searchInput.MinBathrooms != nil {
		query = query.Where("bathrooms >= ?", *searchInput.MinBathrooms)
	}
	if searchInput.MaxBathrooms != nil {
		query = query.Where("bathrooms <= ?", *searchInput.MaxBathrooms)
	}
	if searchInput.MinArea != nil {
		query = query.Where("area >= ?", *searchInput.MinArea)
	}
	if searchInput.MaxArea != nil {
		query = query.Where("area <= ?", *searchInput.MaxArea)
	}
	if searchInput.Type != "" {
		query = query.Where("type = ?", searchInput.Type)
	}
	if searchInput.Available != nil {
		query = query.Where("available = ?", *searchInput.Available)
	}
	if searchInput.Featured != nil {
		query = query.Where("featured = ?", *searchInput.Featured)
	}

	if searchInput.StartDate != nil {
		query = query.Where("start_date >= ?", *searchInput.StartDate)
	}
	if searchInput.EndDate != nil {
		query = query.Where("end_date <= ?", *searchInput.EndDate)
	}

	if searchInput.Amenities != nil {
		query = query.Where("amenities @> ?", searchInput.Amenities)
	}

	if searchInput.OwnerID != nil {
		query = query.Where("owner_id = ?", *searchInput.OwnerID)
	}

	// Paginação
	if searchInput.Limit > 0 {
		query = query.Limit(searchInput.Limit)
	}
	if searchInput.Offset > 0 {
		query = query.Offset(searchInput.Offset)
	}

	err := query.Find(&properties).Error
	if err != nil {
		return []models.Property{}, fmt.Errorf("failed to search properties: %w", err)
	}

	return properties, nil
}

// UpdateProperty atualiza uma propriedade
func UpdateProperty(propertyData models.Property) (models.Property, error) {
	var property models.Property
	err := database.Instance.Where("id = ?", propertyData.ID).First(&property).Error
	if err != nil {
		return models.Property{}, fmt.Errorf("property not found: %w", err)
	}

	// Atualizar campos
	if propertyData.Title != "" {
		property.Title = propertyData.Title
	}
	if propertyData.Description != "" {
		property.Description = propertyData.Description
	}
	if propertyData.Price > 0 {
		property.Price = propertyData.Price
	}
	if propertyData.PriceType != "" {
		property.PriceType = propertyData.PriceType
	}
	if propertyData.Bedrooms > 0 {
		property.Bedrooms = propertyData.Bedrooms
	}
	if propertyData.Bathrooms > 0 {
		property.Bathrooms = propertyData.Bathrooms
	}
	if propertyData.Area > 0 {
		property.Area = propertyData.Area
	}
	if propertyData.Type != "" {
		property.Type = propertyData.Type
	}
	if len(propertyData.Amenities) > 0 {
		property.Amenities = propertyData.Amenities
	}
	if propertyData.Status != "" {
		property.Status = propertyData.Status
	}
	property.Available = propertyData.Available
	property.Featured = propertyData.Featured
	property.Location = propertyData.Location

	err = database.Instance.Save(&property).Error
	if err != nil {
		return models.Property{}, fmt.Errorf("failed to update property: %w", err)
	}

	return property, nil
}

// DeleteProperty deleta uma propriedade
func DeleteProperty(id string) error {
	err := database.Instance.Delete(&models.Property{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete property: %w", err)
	}
	return nil
}

// GetPropertiesByOwner busca propriedades de um proprietário específico
func GetPropertiesByOwner(ownerID uint) ([]models.Property, error) {
	var properties []models.Property
	err := database.Instance.Preload("Images").Where("owner_id = ?", ownerID).Find(&properties).Error
	if err != nil {
		return []models.Property{}, fmt.Errorf("failed to fetch owner properties: %w", err)
	}
	return properties, nil
}

// AddPropertyImage adiciona uma imagem à propriedade
func AddPropertyImage(propertyID uint, imageInput models.PropertyImageInput) (models.PropertyImage, error) {
	image := models.PropertyImage{
		PropertyID: propertyID,
		URL:        imageInput.URL,
		Alt:        imageInput.Alt,
		IsPrimary:  imageInput.IsPrimary,
	}

	// Se esta imagem for primária, remover a primária anterior
	if imageInput.IsPrimary {
		database.Instance.Model(&models.PropertyImage{}).Where("property_id = ? AND is_primary = ?", propertyID, true).Update("is_primary", false)
	}

	err := database.Instance.Create(&image).Error
	if err != nil {
		return models.PropertyImage{}, fmt.Errorf("failed to add property image: %w", err)
	}

	return image, nil
}

// RemovePropertyImage remove uma imagem da propriedade
func RemovePropertyImage(imageID uint) error {
	err := database.Instance.Delete(&models.PropertyImage{}, imageID).Error
	if err != nil {
		return fmt.Errorf("failed to remove property image: %w", err)
	}
	return nil
}

// BookProperty reserva uma propriedade para um período
func BookProperty(propertyID uint, bookingInput models.BookedPeriodInput) (models.BookedPeriod, error) {
	// Verificar se a propriedade está disponível
	var property models.Property
	err := database.Instance.Where("id = ? AND available = ?", propertyID, true).First(&property).Error
	if err != nil {
		return models.BookedPeriod{}, fmt.Errorf("property not available for booking: %w", err)
	}

	// Verificar se o período está disponível
	var existingBooking models.BookedPeriod
	err = database.Instance.Where("property_id = ? AND status IN (?) AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
		propertyID, []string{"pending", "confirmed"}, bookingInput.StartDate, bookingInput.StartDate, bookingInput.EndDate, bookingInput.EndDate, bookingInput.StartDate, bookingInput.EndDate).First(&existingBooking).Error
	if err == nil {
		return models.BookedPeriod{}, fmt.Errorf("property is already booked for this period")
	}

	booking := models.BookedPeriod{
		PropertyID: propertyID,
		StartDate:  bookingInput.StartDate,
		EndDate:    bookingInput.EndDate,
		GuestID:    bookingInput.GuestID,
		Status:     "pending",
	}

	err = database.Instance.Create(&booking).Error
	if err != nil {
		return models.BookedPeriod{}, fmt.Errorf("failed to create booking: %w", err)
	}

	return booking, nil
}

// GetPropertyBookings busca as reservas de uma propriedade
func GetPropertyBookings(propertyID uint) ([]models.BookedPeriod, error) {
	var bookings []models.BookedPeriod
	err := database.Instance.Preload("Guest").Where("property_id = ?", propertyID).Find(&bookings).Error
	if err != nil {
		return []models.BookedPeriod{}, fmt.Errorf("failed to fetch property bookings: %w", err)
	}
	return bookings, nil
}

// UpdateBookingStatus atualiza o status de uma reserva
func UpdateBookingStatus(bookingID uint, status string) (models.BookedPeriod, error) {
	var booking models.BookedPeriod
	err := database.Instance.Where("id = ?", bookingID).First(&booking).Error
	if err != nil {
		return models.BookedPeriod{}, fmt.Errorf("booking not found: %w", err)
	}

	booking.Status = status
	err = database.Instance.Save(&booking).Error
	if err != nil {
		return models.BookedPeriod{}, fmt.Errorf("failed to update booking status: %w", err)
	}

	return booking, nil
}

// GetAvailableProperties busca propriedades disponíveis para um período
func GetAvailableProperties(startDate, endDate time.Time) ([]models.Property, error) {
	var properties []models.Property

	// Subquery para encontrar propriedades com reservas conflitantes
	subQuery := database.Instance.Table("booked_periods").
		Select("DISTINCT property_id").
		Where("status IN (?) AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			[]string{"pending", "confirmed"}, startDate, startDate, endDate, endDate, startDate, endDate)

	err := database.Instance.Preload("Owner").Preload("Images").
		Where("available = ? AND id NOT IN (?)", true, subQuery).
		Find(&properties).Error

	if err != nil {
		return []models.Property{}, fmt.Errorf("failed to fetch available properties: %w", err)
	}

	return properties, nil
}

// UpdatePropertyRating atualiza a avaliação de uma propriedade
func UpdatePropertyRating(propertyID uint, rating float64) error {
	var property models.Property
	err := database.Instance.Where("id = ?", propertyID).First(&property).Error
	if err != nil {
		return fmt.Errorf("property not found: %w", err)
	}

	// Calcular nova média de avaliação

	err = database.Instance.Save(&property).Error
	if err != nil {
		return fmt.Errorf("failed to update property rating: %w", err)
	}

	return nil
}
