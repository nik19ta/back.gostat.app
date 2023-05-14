package database

import (
	"fmt"
	"gostat/pkg/config"
	"gostat/services/auth"
	"gostat/services/links"
	stat "gostat/services/stat"
	"log"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	conf := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.PostgresHost,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDbname,
		conf.PostgresPort,
		conf.PostgresSSLMode,
		conf.PostgresTimezone)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})

	if err != nil {
		log.Panicln(err)
	}

	autoMigrateDB(db)

	return db
}

func autoMigrateDB(db *gorm.DB) {
	// Users
	_ = db.AutoMigrate(&auth.User{})

	// Stats
	_ = db.AutoMigrate(&stat.Visits{})

	// Links
	_ = db.AutoMigrate(&links.Link{})
}
