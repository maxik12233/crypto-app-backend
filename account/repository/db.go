package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbUrl string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//err = db.AutoMigrate(&models.Account{}, &models.Balance{}, &models.Activity{})
	if err != nil {
		return nil, err
	}

	return db, nil
}