package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ktm-m/playground-go-deployment/internal/domain/model"
)

type UserPort interface {
	RegisterRoutes()
	GetUserByID(c fiber.Ctx) error
}

type user struct {
	app *fiber.App
}

func newUser(app *fiber.App) UserPort {
	return &user{
		app: app,
	}
}

func (u *user) RegisterRoutes() {
	u.app.Post("/api/user/id/inquiry", u.GetUserByID)
}

func (u *user) GetUserByID(c fiber.Ctx) error {
	return c.JSON(model.GetUserByIDRespContent{
		ID:          "1",
		FirstnameTH: "ทดสอบ",
		LastnameTH:  "ทดสอบ",
		Age:         20,
	})
}
