package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"
)

type (
	MySQL struct {
		Db *sql.DB
	}
)

func NewMySQLConnector(cfg *configs.Config) *MySQL {
	cfgDb := cfg.Database

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfgDb.User, cfgDb.Password, cfgDb.Host, cfgDb.Port, cfgDb.Name)
	db, err := sql.Open(cfgDb.Driver, dataSource)
	if err != nil {
		panic(err)
	}

	createTables(db)

	return &MySQL{Db: db}
}

func createTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id CHAR(36) PRIMARY KEY,
		price DOUBLE NOT NULL,
		tax DOUBLE NOT NULL,
		final_price DOUBLE NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("error creating table 'orders': %s", err))
	}
}
