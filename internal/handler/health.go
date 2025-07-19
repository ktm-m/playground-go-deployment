package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ktm-m/playground-go-deployment/internal/domain/constant"
	"github.com/ktm-m/playground-go-deployment/internal/domain/model"
)

type HealthPort interface {
	RegisterRoutes()
	HealthCheck(c fiber.Ctx) error
}

type health struct {
	app *fiber.App
}

func newHealth(app *fiber.App) HealthPort {
	return &health{
		app: app,
	}
}

func (h *health) RegisterRoutes() {
	h.app.Get("/api/health", h.HealthCheck)
}

func (h *health) HealthCheck(c fiber.Ctx) error {
	return c.JSON(model.HealthCheckRespContent{
		Status: constant.HealthCheckStatusOK,
	})
}
