package vehicle

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"kukuhkkh.id/learn/bengkel/domain"
	"kukuhkkh.id/learn/bengkel/internal/util"
	"time"
)

type api struct {
	vehicleService domain.VehicleService
}

func NewApi(app *fiber.App, vehicleService domain.VehicleService) {
	api := api{
		vehicleService,
	}

	app.Get("/v1/vehicle-histories", api.GetVehicleHistories)
	app.Post("/v1/vehicle-histories", api.StoreVehicleHistory)
}

func (a api) GetVehicleHistories(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	vin := ctx.Query("vin")
	if vin == "" {
		apiResponse := domain.ApiResponse{
			Code:    "400",
			Message: "VIN is required",
			Data:    nil,
		}
		util.ResponseInterceptor(c, &apiResponse)

		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.FindHistorical(c, vin)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}

func (a api) StoreVehicleHistory(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var req domain.VehicleHistoricalRequest
	if err := ctx.BodyParser(&req); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "400",
			Message: "Bad Request",
			Data:    nil,
		}
		util.ResponseInterceptor(c, &apiResponse)

		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.StoreHistorical(c, req)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}
