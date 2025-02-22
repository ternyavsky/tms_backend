package db

import (
	"fmt"
	"time"

	"github.com/ternyavsky/tms_backend/config"
	"github.com/ternyavsky/tms_backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabse() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", config.GlobalConfig.DBHost, config.GlobalConfig.DBUser, config.GlobalConfig.DBPassword, config.GlobalConfig.DBName, config.GlobalConfig.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Europe/Moscow")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		fmt.Println("error this")
		return err
	}
	db.AutoMigrate(
		&domain.Link{},
		&domain.CallbackOrder{},
		&domain.SourceUpdate{},
		&domain.SourcePost{},
		&domain.SourceStatistic{},
		&domain.Session{},
		&domain.PostStatistic{},
		&domain.Source{},
	)
	return nil
}
