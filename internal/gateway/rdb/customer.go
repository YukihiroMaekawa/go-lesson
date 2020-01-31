package rdb

import (
	"database/sql"
	"time"

	"github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
	_ "github.com/lib/pq"
	"go.uber.org/multierr"
)

type App struct {
	db *sql.DB
}

func NewCustomer(db *sql.DB) App {
	return App{
		db: db,
	}
}

func (app App) CreateCustomer(request *entity.RegisterCustomerRequest) (string, error) {
	tx, _ := app.db.Begin()

	t := time.Now()
	_, err := tx.Exec(`INSERT INTO customers 
							(name ,address1 ,address2 ,address3 ,created_at ,updated_at) 
				 	 VALUES ($1 ,$2 ,$3 ,$4 ,$5 ,$6)`,
		request.Name, request.Address1, request.Address2, request.Address3, t, t,
	)
	if err != nil {
		return "", multierr.Append(err, tx.Rollback())
	}

	err = tx.Commit()
	if err != nil {
		return "", multierr.Append(err, tx.Rollback())
	}
	return "customer:001", nil
}
