package handler

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Input struct {
	Operator string  `json:"operator"`
	Number_1 float64 `json:"num1"`
	Number_2 float64 `json:"num2"`
}

type CheckOperator func(input Input) (result float64, err error)

type GetRedis func(ctx context.Context, key string) (result float64, err error)

type SetRedis func(ctx context.Context, key string, result float64)

func NewHandler(CheckOperator CheckOperator, GetRedis GetRedis, SetRedis SetRedis) fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		input := Input{}
		err := c.BodyParser(&input)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "BodyParser error")
		}

		key := fmt.Sprintf("calculate::%v%v%v", input.Number_1, input.Operator, input.Number_2)

		if result, err := GetRedis(c.Context(), key); err == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"result": result,
			})
		}

		result, err := CheckOperator(input)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		SetRedis(c.Context(), key, result)
		
		fmt.Println(".....init set.....")
		return c.Status(200).JSON(fiber.Map{
			"result": result,
		})
	}
}
