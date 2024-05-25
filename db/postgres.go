package db

import (
	"fmt"
	"go_eduhub/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{}

func NewDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connect(creds *model.Credential) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
