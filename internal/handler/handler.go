package handler

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	Health HealthPort
	User   UserPort
}

func NewHandler(app *fiber.App) *Handler {
	return &Handler{
		Health: newHealth(app),
		User:   newUser(app),
	}
}

func (h *Handler) RegisterRoutes() {
	h.Health.RegisterRoutes()
	h.User.RegisterRoutes()
}
