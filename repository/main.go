package repository

import (
	"assignment-2/configs"
	pgsql "assignment-2/helpers/database"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	dbSql *sql.DB
)

func init() {
	config := configs.InitConfig()
	pgsql.StartDB(config.Database)

	db = pgsql.GetDB()
	dbSql, _ = db.DB()
	dbSql.SetMaxIdleConns(5)
	dbSql.SetMaxOpenConns(1000)
	dbSql.SetConnMaxLifetime(time.Hour)
}
