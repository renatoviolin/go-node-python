package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/renatoviolin/go-node/dto"
	"github.com/renatoviolin/go-node/usecase"
)

type FiberHandler struct {
	fiber *fiber.App
	uc    *usecase.FindAll
}

func NewHandler(uc *usecase.FindAll) *FiberHandler {
	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	return &FiberHandler{
		fiber: fiber,
		uc:    uc,
	}
}

func (h *FiberHandler) SetupAndRun(address string) {
	h.fiber.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	h.fiber.Get("/single_json", func(c *fiber.Ctx) error {
		return c.JSON(dto.NewPayload())
	})

	h.fiber.Get("/mongo_json", func(c *fiber.Ctx) error {
		result, err := h.uc.Execute()
		if err != nil {
			return c.JSON(err.Error())
		}
		return c.JSON(result)
	})

	fmt.Printf("Fiber running at %s\n", address)
	h.fiber.Listen(address)
}
