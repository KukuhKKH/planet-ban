package domain

import (
	"context"
	"time"
)

type History struct {
	ID        int       `db:"id"`
	NoRangka  string    `db:"no_rangka"`
	Merek     string    `db:"merek"`
	UpdatedAt time.Time `db:"updated_at"`
}

type HistoryDetail struct {
	ID         int       `db:"id"`
	HistoryID  int       `db:"history_id"`
	CustomerID int       `db:"customer_id"`
	PIC        string    `db:"pic"`
	PlatNomor  string    `db:"plat_nomor"`
	Notes      string    `db:"notes"`
	Date       time.Time `db:"date"`
}

type HistoryRepository interface {
	FindByID(ctx context.Context, id int) (History, error)
	FindByNoRangka(ctx context.Context, no string) (History, error)
	FindDetailByHistory(ctx context.Context, id int) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *History) error
	InsertDetail(ctx context.Context, detail *HistoryDetail) error
}

type HistoryService interface {
}
