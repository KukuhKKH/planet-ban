package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"kukuhkkh.id/learn/bengkel/internal/component"
	"kukuhkkh.id/learn/bengkel/internal/config"
	"kukuhkkh.id/learn/bengkel/internal/modul/customer"
	"kukuhkkh.id/learn/bengkel/internal/modul/history"
	"kukuhkkh.id/learn/bengkel/internal/modul/vehicle"
)

func main() {
	conf := config.Get()
	conn := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(conn)
	vehicleRespository := vehicle.NewRepository(conn)
	historyRespository := history.NewRepository(conn)

	customerService := customer.NewService(customerRepository)
	vehicleService := vehicle.NewService(vehicleRespository, historyRespository)

	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} ${latency} - ${status} ${method} ${path}\n",
	}))

	customer.NewApi(app, customerService)
	vehicle.NewApi(app, vehicleService)

	_ = app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
