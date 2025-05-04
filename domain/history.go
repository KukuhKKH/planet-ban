package domain

import (
	"context"
	"time"
)

type HistoryDetail struct {
	ID          int       `db:"id"`
	VehicleID   int       `db:"vehicle_id"`
	CustomerID  int       `db:"customer_id"`
	PIC         string    `db:"pic"`
	PlateNumber string    `db:"plate_number"`
	Notes       string    `db:"notes"`
	Date        time.Time `db:"date"`
}

type HistorycalData struct {
	VehicleID   int    `json:"vehicle_id"`
	CustomerID  int    `json:"customer_id"`
	PIC         string `json:"pic"`
	PlateNumber string `json:"plate_number"`
	Notes       string `json:"notes"`
	ComeAt      string `json:"come_at"`
}

type HistoryRepository interface {
	FindDetailByVehicle(ctx context.Context, id int) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *HistoryDetail) error
}

type HistoryService interface {
}
