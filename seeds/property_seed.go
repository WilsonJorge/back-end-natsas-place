package seeds

import (
	"encoding/json"
	"log"
	"natsas-place/database"
	"natsas-place/models"

	"golang.org/x/crypto/bcrypt"
)

// PropertySeedFixed cria propriedades de exemplo com correções
func PropertySeedFixed() {
	log.Println("🌱 Starting fixed property seed...")

	// Criar usuários de exemplo primeiro
	users := createSampleUsersFixed()

	// Criar propriedades de exemplo sem amenities
	properties := createSamplePropertiesFixed(users)

	// Adicionar amenities às propriedades após a criação
	addAmenitiesToPropertiesFixed(properties)

	// Adicionar imagens às propriedades
	addSampleImagesFixed(properties)

	log.Println("✅ Fixed property seed completed successfully!")
}

// createSampleUsersFixed cria usuários de exemplo com senhas
func createSampleUsersFixed() []models.User {
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("❌ Error generating password: %v", err)
	}

	users := []models.User{
		{
			Name:         "João Silva",
			Email:        "joao.silva@example.com",
			Password:     string(password),
			Phone:        "+55 11 99999-9999",
			Role:         "user",
			IsVerified:   true,
			IsActive:     true,
			ProfileImage: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&h=150&fit=crop&crop=face",
		},
		{
			Name:         "Maria Santos",
			Email:        "maria.santos@example.com",
			Password:     string(password),
			Phone:        "+55 11 88888-8888",
			Role:         "user",
			IsVerified:   true,
			IsActive:     true,
			ProfileImage: "https://images.unsplash.com/photo-1494790108755-2616b612b786?w=150&h=150&fit=crop&crop=face",
		},
		{
			Name:         "Pedro Costa",
			Email:        "pedro.costa@example.com",
			Password:     string(password),
			Phone:        "+55 11 77777-7777",
			Role:         "user",
			IsVerified:   true,
			IsActive:     true,
			ProfileImage: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=150&h=150&fit=crop&crop=face",
		},
		{
			Name:         "Ana Oliveira",
			Email:        "ana.oliveira@example.com",
			Password:     string(password),
			Phone:        "+55 11 66666-6666",
			Role:         "user",
			IsVerified:   true,
			IsActive:     true,
			ProfileImage: "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=150&h=150&fit=crop&crop=face",
		},
	}

	for i := range users {
		// Verificar se o usuário já existe
		var existingUser models.User
		err := database.Instance.Where("email = ?", users[i].Email).First(&existingUser).Error
		if err != nil {
			// Usuário não existe, criar
			err = database.Instance.Create(&users[i]).Error
			if err != nil {
				log.Printf("❌ Error creating user %s: %v", users[i].Email, err)
			} else {
				log.Printf("✅ Created user: %s", users[i].Email)
			}
		} else {
			// Usuário já existe, usar o existente
			users[i] = existingUser
			log.Printf("ℹ️  User already exists: %s", users[i].Email)
		}
	}

	return users
}

// createSamplePropertiesFixed cria propriedades de exemplo sem amenities
func createSamplePropertiesFixed(users []models.User) []models.Property {
	properties := []models.Property{
		{
			Title:       "Casa Moderna na Vila Madalena",
			Description: "Linda casa moderna com 3 quartos, 2 banheiros e uma vista incrível da cidade. Localizada em uma das áreas mais desejadas de São Paulo, próxima a restaurantes, bares e transporte público.",
			Price:       450.0,
			Location: models.Location{
				Address: "Rua Harmonia, 123",
				City:    "São Paulo",
				State:   "SP",
				ZipCode: "05435-000",
				Coordinates: struct {
					Lat float64 `json:"lat" gorm:"not null" validate:"required"`
					Lng float64 `json:"lng" gorm:"not null" validate:"required"`
				}{
					Lat: -23.5505,
					Lng: -46.6333,
				},
			},
			PriceType: "daily",
			Bedrooms:  3,
			Bathrooms: 2,
			Area:      120.0,
			Type:      "house",
			OwnerID:   users[0].ID,
			Status:    "available",
			Available: true,
			Featured:  true,
		},
		{
			Title:       "Apartamento Elegante no Jardins",
			Description: "Apartamento luxuoso com acabamento de alto padrão, localizado no coração do Jardins. Ideal para quem busca conforto e sofisticação em uma das áreas mais nobres de São Paulo.",
			Price:       380.0,
			Location: models.Location{
				Address: "Alameda Santos, 456",
				City:    "São Paulo",
				State:   "SP",
				ZipCode: "01418-000",
				Coordinates: struct {
					Lat float64 `json:"lat" gorm:"not null" validate:"required"`
					Lng float64 `json:"lng" gorm:"not null" validate:"required"`
				}{
					Lat: -23.5686,
					Lng: -46.6654,
				},
			},
			PriceType: "daily",
			Bedrooms:  2,
			Bathrooms: 2,
			Area:      85.0,
			Type:      "apartment",
			OwnerID:   users[1].ID,
			Status:    "available",
			Available: true,
			Featured:  true,
		},
		{
			Title:       "Studio Acolhedor na Consolação",
			Description: "Studio moderno e funcional, perfeito para uma estadia confortável. Localizado próximo ao metrô e com fácil acesso a todas as atrações da cidade.",
			Price:       180.0,
			Location: models.Location{
				Address: "Rua da Consolação, 789",
				City:    "São Paulo",
				State:   "SP",
				ZipCode: "01302-001",
				Coordinates: struct {
					Lat float64 `json:"lat" gorm:"not null" validate:"required"`
					Lng float64 `json:"lng" gorm:"not null" validate:"required"`
				}{
					Lat: -23.5505,
					Lng: -46.6333,
				},
			},
			PriceType: "daily",
			Bedrooms:  1,
			Bathrooms: 1,
			Area:      45.0,
			Type:      "studio",
			OwnerID:   users[2].ID,
			Status:    "available",
			Available: true,
			Featured:  false,
		},
		{
			Title:       "Casa de Campo em Atibaia",
			Description: "Casa de campo com vista para a serra, ideal para quem busca tranquilidade e contato com a natureza. Perfeita para fins de semana ou férias em família.",
			Price:       320.0,
			Location: models.Location{
				Address: "Estrada Municipal, 321",
				City:    "Atibaia",
				State:   "SP",
				ZipCode: "12940-000",
				Coordinates: struct {
					Lat float64 `json:"lat" gorm:"not null" validate:"required"`
					Lng float64 `json:"lng" gorm:"not null" validate:"required"`
				}{
					Lat: -23.1169,
					Lng: -46.5505,
				},
			},
			PriceType: "daily",
			Bedrooms:  4,
			Bathrooms: 3,
			Area:      200.0,
			Type:      "house",
			OwnerID:   users[3].ID,
			Status:    "available",
			Available: true,
			Featured:  true,
		},
		{
			Title:       "Loft Industrial no Bom Retiro",
			Description: "Loft com design industrial, localizado em um bairro em transformação. Ideal para quem busca um espaço único e autêntico.",
			Price:       250.0,
			Location: models.Location{
				Address: "Rua Vitorino Carmilo, 654",
				City:    "São Paulo",
				State:   "SP",
				ZipCode: "01153-000",
				Coordinates: struct {
					Lat float64 `json:"lat" gorm:"not null" validate:"required"`
					Lng float64 `json:"lng" gorm:"not null" validate:"required"`
				}{
					Lat: -23.5200,
					Lng: -46.6333,
				},
			},
			PriceType: "daily",
			Bedrooms:  2,
			Bathrooms: 1,
			Area:      75.0,
			Type:      "apartment",
			OwnerID:   users[0].ID,
			Status:    "available",
			Available: true,
			Featured:  false,
		},
		{
			Title:       "Casa na Praia em Ubatuba",
			Description: "Casa na praia com vista para o mar, localizada em uma das praias mais bonitas de Ubatuba. Ideal para férias em família ou casais.",
			Price:       280.0,
			Location: models.Location{
				Address: "Rua das Palmeiras, 987",
				City:    "Ubatuba",
				State:   "SP",
				ZipCode: "11680-000",
				Coordinates: struct {
					Lat float64 `json:"lat" gorm:"not null" validate:"required"`
					Lng float64 `json:"lng" gorm:"not null" validate:"required"`
				}{
					Lat: -23.4339,
					Lng: -45.0839,
				},
			},
			PriceType: "daily",
			Bedrooms:  3,
			Bathrooms: 2,
			Area:      150.0,
			Type:      "house",
			OwnerID:   users[1].ID,
			Status:    "available",
			Available: true,
			Featured:  true,
		},
	}

	for i := range properties {
		// Verificar se a propriedade já existe
		var existingProperty models.Property
		err := database.Instance.Where("title = ? AND owner_id = ?", properties[i].Title, properties[i].OwnerID).First(&existingProperty).Error
		if err != nil {
			// Propriedade não existe, criar
			err = database.Instance.Create(&properties[i]).Error
			if err != nil {
				log.Printf("❌ Error creating property %s: %v", properties[i].Title, err)
			} else {
				log.Printf("✅ Created property: %s", properties[i].Title)
			}
		} else {
			// Propriedade já existe, usar a existente
			properties[i] = existingProperty
			log.Printf("ℹ️  Property already exists: %s", properties[i].Title)
		}
	}

	return properties
}

// addAmenitiesToPropertiesFixed adiciona amenities às propriedades após a criação
func addAmenitiesToPropertiesFixed(properties []models.Property) {
	amenitiesMap := map[string][]string{
		"Casa Moderna na Vila Madalena":   {"WiFi", "Ar condicionado", "Cozinha equipada", "Estacionamento", "Churrasqueira", "Piscina"},
		"Apartamento Elegante no Jardins": {"WiFi", "Ar condicionado", "Cozinha equipada", "Portaria 24h", "Academia", "Salão de festas"},
		"Studio Acolhedor na Consolação":  {"WiFi", "Ar condicionado", "Cozinha compacta", "TV", "Roupas de cama"},
		"Casa de Campo em Atibaia":        {"WiFi", "Lareira", "Churrasqueira", "Piscina", "Quadra de tênis", "Trilhas", "Vista para a serra"},
		"Loft Industrial no Bom Retiro":   {"WiFi", "Ar condicionado", "Cozinha equipada", "Design industrial", "Área de trabalho"},
		"Casa na Praia em Ubatuba":        {"WiFi", "Vista para o mar", "Churrasqueira", "Piscina", "Acesso direto à praia", "Estacionamento"},
	}

	for _, property := range properties {
		amenities, exists := amenitiesMap[property.Title]
		if !exists {
			continue
		}

		// Converter amenities para JSON
		amenitiesJSON, err := json.Marshal(amenities)
		if err != nil {
			log.Printf("❌ Error marshaling amenities for property %s: %v", property.Title, err)
			continue
		}

		// Atualizar a propriedade com as amenities
		err = database.Instance.Model(&property).Update("amenities", amenitiesJSON).Error
		if err != nil {
			log.Printf("❌ Error updating amenities for property %s: %v", property.Title, err)
		} else {
			log.Printf("✅ Updated amenities for property: %s", property.Title)
		}
	}
}

// addSampleImagesFixed adiciona imagens de exemplo às propriedades
func addSampleImagesFixed(properties []models.Property) {
	// Imagens para cada propriedade
	propertyImages := map[string][]string{
		"Casa Moderna na Vila Madalena": {
			"https://images.unsplash.com/photo-1564013799919-ab600027ffc6?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1571896349842-33c89424de2d?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1560448204-e02f11c3d0e2?w=800&h=600&fit=crop",
		},
		"Apartamento Elegante no Jardins": {
			"https://images.unsplash.com/photo-1502672260266-1c1ef2d93688?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1493809842364-78817add7ffb?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1502005229762-cf1b2da7c5d6?w=800&h=600&fit=crop",
		},
		"Studio Acolhedor na Consolação": {
			"https://images.unsplash.com/photo-1522708323590-d24dbb6b0267?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1502005229762-cf1b2da7c5d6?w=800&h=600&fit=crop",
		},
		"Casa de Campo em Atibaia": {
			"https://images.unsplash.com/photo-1564013799919-ab600027ffc6?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1571896349842-33c89424de2d?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1560448204-e02f11c3d0e2?w=800&h=600&fit=crop",
		},
		"Loft Industrial no Bom Retiro": {
			"https://images.unsplash.com/photo-1502672260266-1c1ef2d93688?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1493809842364-78817add7ffb?w=800&h=600&fit=crop",
		},
		"Casa na Praia em Ubatuba": {
			"https://images.unsplash.com/photo-1564013799919-ab600027ffc6?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1571896349842-33c89424de2d?w=800&h=600&fit=crop",
			"https://images.unsplash.com/photo-1560448204-e02f11c3d0e2?w=800&h=600&fit=crop",
		},
	}

	for _, property := range properties {
		images, exists := propertyImages[property.Title]
		if !exists {
			continue
		}

		// Verificar se a propriedade já tem imagens
		var existingImages []models.PropertyImage
		err := database.Instance.Where("property_id = ?", property.ID).Find(&existingImages).Error
		if err == nil && len(existingImages) > 0 {
			log.Printf("ℹ️  Property %s already has images", property.Title)
			continue
		}

		// Adicionar imagens
		for i, imageURL := range images {
			image := models.PropertyImage{
				PropertyID: property.ID,
				URL:        imageURL,
				Alt:        property.Title,
				IsPrimary:  i == 0, // Primeira imagem é a primária
			}

			err := database.Instance.Create(&image).Error
			if err != nil {
				log.Printf("❌ Error adding image to property %s: %v", property.Title, err)
			} else {
				log.Printf("✅ Added image %d to property: %s", i+1, property.Title)
			}
		}
	}
}
