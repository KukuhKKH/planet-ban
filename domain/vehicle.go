package domain

import (
	"context"
	"time"
)

type Vehicle struct {
	ID        int       `db:"id"`
	VIN       string    `db:"vin"`
	Brand     string    `db:"brand"`
	UpdatedAt time.Time `db:"updated_at"`
}

type VehicleHistorical struct {
	ID        int              `json:"id"`
	VIN       string           `json:"vin"`
	Brand     string           `json:"brand"`
	Histories []HistorycalData `json:"histories"`
}

type VehicleRepository interface {
	FindByID(ctx context.Context, id int) (Vehicle, error)
	FindByVIN(ctx context.Context, vin string) (Vehicle, error)
	Insert(ctx context.Context, vehicle *Vehicle) error
}

type VehicleService interface {
	FindHistorical(ctx context.Context, vin string) ApiResponse
	StoreHistorical(ctx context.Context, request VehicleHistoricalRequest) ApiResponse
}

type VehicleHistoricalRequest struct {
	CustomerID  int    `json:"customer_id"`
	VIN         string `json:"vin"`
	Brand       string `json:"brand"`
	PIC         string `json:"pic"`
	PlateNumber string `json:"plate_number"`
	Notes       string `json:"notes"`
}
