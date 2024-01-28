package main

import (
	"fmt"

	"github.com/LautaroNSantillan/my-go-journey/tree/crm-fiber/database"
	"github.com/LautaroNSantillan/my-go-journey/tree/crm-fiber/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func main() {
	app := fiber.New()
	initDB()
	setUpRoutes(app)
	app.Listen(8000)
	defer database.DBCon.Close()

}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDB() {
	var err error
	database.DBCon, err := gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("Connected to DB")
	database.DBCon.AutoMigrate(&lead.Lead{})
	fmt.Println("DB Migrated")

}
