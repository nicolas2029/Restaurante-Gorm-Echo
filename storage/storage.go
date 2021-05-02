package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open("postgres://admin_restaurant:RestAuraNt_pgsql.561965697@localhost:5433/restaurante_gorm_echo?sslmode=disable"))
		if err != nil {
			log.Fatalf("no se pudo abrir la base de datos: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
