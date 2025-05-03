package domain

import "time"

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
}

type HistoryService interface {
}
