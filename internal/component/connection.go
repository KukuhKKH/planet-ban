package component

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"kukuhkkh.id/learn/bengkel/internal/config"
	"log"
)

func GetDatabaseConnection(conf config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.DB.User,
		conf.DB.Pass,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Name,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening database connection: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err.Error())
	}

	return db
}
