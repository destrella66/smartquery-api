package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		"roundhouse.proxy.rlwy.net",
		"postgres",
		"AFd5Cec6-D232*aaaFDB-e4DF6fgfbDC",
		"railway",
		"12542")

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	automigrate(db)

	return db, err
}

func automigrate(db *gorm.DB) {
	db.Debug().AutoMigrate(
		Chamado{},
		Admin{},
		Cliente{},
		Produto{},
	)
}
