package history

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"kukuhkkh.id/learn/bengkel/domain"
	"time"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{
		db: goqu.Dialect("mysql").DB(con),
	}
}

func (r repository) FindDetailByVehicle(ctx context.Context, id int) (result []domain.HistoryDetail, err error) {
	dataset := r.db.From("history_details").Where(goqu.Ex{
		"vehicle_id": id,
	}).Order(goqu.I("id").Asc())

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (r repository) Insert(ctx context.Context, history *domain.HistoryDetail) error {
	history.Date = time.Now()

	exceutor := r.db.Insert("history_details").Rows(goqu.Record{
		"vehicle_id":   history.VehicleID,
		"customer_id":  history.CustomerID,
		"pic":          history.PIC,
		"plate_number": history.PlateNumber,
		"notes":        history.Notes,
		"date":         history.Date,
	}).Executor()

	res, err := exceutor.ExecContext(ctx)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	history.ID = int(id)
	return nil
}
