package lead

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/LautaroNSantillan/my-go-journey/tree/crm-fiber/database"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(ctx *fiber.Ctx) {
	db := database.DBCon
	var leads []Lead

	db.Find(&leads)

	ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBCon
	var lead Lead

	db.Find(&lead, id)

	ctx.JSON(lead)
}

func NewLead(ctx *fiber.Ctx) {
	db := database.DBCon
	lead := new(Lead)

	if err := ctx.BodyParser(lead); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) {
	db := database.DBCon
	id := ctx.Params("id")
	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		ctx.Status(500).Send("No lead found")
		return
	}

	db.Delete(&lead)
	ctx.Send(fmt.Printf("Deleted Lead with ID: %v", id))
}
