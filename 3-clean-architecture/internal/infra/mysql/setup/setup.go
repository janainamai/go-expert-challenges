package setup

import (
	"database/sql"
	"fmt"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type (
	MySQL struct {
		DB *sql.DB
	}
)

func NewMySQLConnector(cfg *configs.Config) *MySQL {
	cfgDB := cfg.Database

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfgDB.User, cfgDB.Password, cfgDB.Host, cfgDB.Port, cfgDB.Name)
	db, err := sql.Open(cfgDB.Driver, dataSource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Erro ao verificar a conexão: %s", err))
	}

	runMigrations(cfg)

	return &MySQL{DB: db}
}

func runMigrations(cfg *configs.Config) {
	cfgDB := cfg.Database

	dataSource := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s",
		cfgDB.Driver, cfgDB.User, cfgDB.Password, cfgDB.Host, cfgDB.Port, cfgDB.Name)

	m, err := migrate.New(
		"file://migrations/sql",
		dataSource)
	if err != nil {
		panic(fmt.Sprintf("error initing migrate configs: %s", err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Sprintf("error executing migrations: %s", err))
	}
}
