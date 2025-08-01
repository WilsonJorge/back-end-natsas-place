package seeds

import (
	"log"
	"natsas-place/database"
)

// RunSeeds executa todos os seeds
func RunSeeds() {
	log.Println("🚀 Starting database seeds...")

	// Conectar ao banco de dados
	database.ConnectBD()

	// Executar migrações

	// Executar seeds
	PropertySeedFixed()

	log.Println("✅ All seeds completed successfully!")
}
