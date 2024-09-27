package controller

import (
	"github.com/B6137151/GDZ-Commerce/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StoreController struct {
	storeService service.StoreService
}

func NewStoreController(storeService service.StoreService) *StoreController {
	return &StoreController{storeService: storeService}
}

func (c *StoreController) CreateStore(ctx *fiber.Ctx) error {
	var input struct {
		StoreName string `json:"store_name"`
		Location  string `json:"location"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	store, err := c.storeService.CreateStore(input.StoreName, input.Location)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create store"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(store)
}

func (c *StoreController) GetStore(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid store ID"})
	}

	store, err := c.storeService.GetStore(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Store not found"})
	}

	return ctx.JSON(store)
}

func (c *StoreController) UpdateStore(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid store ID"})
	}

	var input struct {
		StoreName string `json:"store_name"`
		Location  string `json:"location"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	store, err := c.storeService.UpdateStore(id, input.StoreName, input.Location)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update store"})
	}

	return ctx.JSON(store)
}

func (c *StoreController) DeleteStore(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid store ID"})
	}

	if err := c.storeService.DeleteStore(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete store"})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
