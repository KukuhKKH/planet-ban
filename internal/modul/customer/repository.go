package customer

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

func (r repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Order(goqu.I("name").Asc())

	err = dataset.ScanStructsContext(ctx, &customers)
	if err != nil {
		return nil, err
	}

	return
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{
		db: goqu.Dialect("mysql").DB(con),
	}
}

func (r repository) FindByID(ctx context.Context, id int) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &customer)
	if err != nil {
		return domain.Customer{}, err
	}

	return
}

func (r repository) FindByIds(ctx context.Context, ids []int) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": ids,
	})

	err = dataset.ScanStructsContext(ctx, &customers)
	if err != nil {
		return nil, err
	}

	return
}

func (r repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"phone": phone,
	})

	_, err = dataset.ScanStructContext(ctx, &customer)
	if err != nil {
		return domain.Customer{}, err
	}

	return
}

func (r repository) Insert(ctx context.Context, customer *domain.Customer) error {
	executor := r.db.Insert("customers").Rows(goqu.Record{
		"name":       customer.Name,
		"phone":      customer.Phone,
		"created_at": customer.CreatedAt,
	}).Executor()

	res, err := executor.ExecContext(ctx)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	customer.ID = int(id)
	return err
}
