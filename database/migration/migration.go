package migration

import (
	"belajar_golang_fiber/database"
	"belajar_golang_fiber/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated")
}