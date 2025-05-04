package vehicle

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"kukuhkkh.id/learn/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{
		db: goqu.Dialect("mysql").DB(con),
	}
}

func (r repository) FindByID(ctx context.Context, id int) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

func (r repository) FindByVIN(ctx context.Context, vin string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicle").Where(goqu.Ex{
		"vin": vin,
	})

	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

func (r repository) Insert(ctx context.Context, vehicle *domain.Vehicle) error {
	executor := r.db.Insert("vehicle").Rows(goqu.Record{
		"vin":        vehicle.VIN,
		"brand":      vehicle.Brand,
		"updated_at": vehicle.UpdatedAt,
	}).Executor()

	res, err := executor.ExecContext(ctx)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	vehicle.ID = int(id)
	return nil
}
