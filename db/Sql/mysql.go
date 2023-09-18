package sql

import (
	"fmt"

	"github.com/aminkhn/portfolio/config"
	"github.com/aminkhn/portfolio/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DataBase *sqlx.DB

type Queries struct {
	*service.UserQuery
	*service.PostQuery
	*service.PortfolioQuery
	*service.ManagementQuery
}

func MysqlConnect() (*Queries, error) {
	// load env configs
	env, _ := config.LoadConfig()
	// establish mysql connection
	Dsn := env.DBUserame + ":" + env.DBPassword + "@tcp(" + env.DBHost + ":" + env.DBPort + ")/" + env.DBName + "?" + "parseTime=true"
	DataBase, err := sqlx.Connect("mysql", Dsn)
	if err != nil {
		return nil, err
	}
	// ping connection check
	if err := DataBase.Ping(); err != nil {
		defer DataBase.Close()
		err = fmt.Errorf("error, ping did not send to database, %w", err)
		return nil, err
	}
	// service queries
	return &Queries{
		UserQuery: &service.UserQuery{DB: DataBase},
		PostQuery: &service.PostQuery{DB: DataBase},
	}, nil
}

// close mysql connection
func MysqlClose() {
	if DataBase != nil {
		sqlDB := DataBase.DB
		err := sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}
}
