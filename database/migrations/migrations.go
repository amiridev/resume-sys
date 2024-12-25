package migrations

import (
	"fmt"
	"log"
	"resume-sys/core/config"
	"resume-sys/database"
	"resume-sys/database/models"
)

func Migrate() {
	config.InitEnvs()

	database.InitAndConnect()
	db := database.Connection()

	err := db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Skill{},
	)

	if err != nil {
		log.Fatal("Cannot migrate")
		return
	}

	fmt.Println("Migration done ..")
}