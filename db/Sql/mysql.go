package sql

import (
	"github.com/aminkhn/portfolio/config"
	"github.com/aminkhn/portfolio/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err      error
	DataBase *gorm.DB
)

func MysqlConnect() {
	env, _ := config.LoadConfig()

	Dsn := env.DBUserame + ":" + env.DBPassword + "@tcp(" + env.DBHost + ":" + env.DBPort + ")/" + env.DBName + "?" + "parseTime=true"
	DataBase, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DataBase.InstanceSet("gorm:table_options", "ENGINE="+env.DBENGINE)
	sqlDB, _ := DataBase.DB()
	sqlDB.SetMaxIdleConns(10)
	// migration models to database
	DataBase.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.WebSite{},
		&model.Portfolio{},
		&model.Project{},
		&model.Technology{},
	)
}

// close mysql connection
func MysqlClose() {
	if DataBase != nil {
		sqlDB, _ := DataBase.DB()
		err := sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}
}
