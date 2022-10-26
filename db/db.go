package db

import (
	"examples/identity/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(s *config.Specification) *gorm.DB {
	sqlDB := mysql.Open(s.DBConnection)
	db, _ := gorm.Open(sqlDB, &gorm.Config{})
	return db
}
