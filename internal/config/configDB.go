package config

import (
	"fmt"
	"log"

	"go000/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("Chatdb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar la DB", err)
	}

	fmt.Println("Conexion a la DB exitosa")
}

// Migracion de los modelos
func MigrateModels() error {
	modelsList := []interface{}{
		&models.User{},
		&models.Post{},
	}

	for _, model := range modelsList {
		if err := DB.AutoMigrate(model); err != nil {
			log.Fatal("Error al realizar la migracion de los modelos:", err)
		}
	}
	fmt.Println("Migraciones realizadas con exito")
	return nil
}
