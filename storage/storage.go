package storage

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbUser struct {
	TypeDB   string `json:"type_db"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	NameDB   string `json:"name_db"`
	Host     string `json:"host"`
}

var (
	db   *gorm.DB
	once sync.Once
)

// Drivers
const (
	Postgres string = "POSTGRES"
	MySql    string = "MYSQL"
	Oracle   string = "ORACLE"
)

// New create the connection with db
func New() {
	once.Do(func() {
		u := loadData()
		switch u.TypeDB {
		case Postgres:
			newPostgresDB(&u)
		case MySql:
			newMySqlBD(&u)
		}
	})
}

func getEnv(env string) (string, error) {
	s, f := os.LookupEnv(env)
	if !f {
		return "", fmt.Errorf("environment variable (%s) not found", env)
	}
	return s, nil
}

func loadData() dbUser {
	typeDb, err := getEnv("RGE_TYPE")
	if err != nil {
		log.Fatalf(err.Error())
	}
	user, err := getEnv("RGE_USER")
	if err != nil {
		log.Fatalf(err.Error())
	}
	password, err := getEnv("RGE_PASSWORD")
	if err != nil {
		log.Fatalf(err.Error())
	}
	port, err := getEnv("RGE_PORT")
	if err != nil {
		log.Fatalf(err.Error())
	}
	name, err := getEnv("RGE_NAME_DB")
	if err != nil {
		log.Fatalf(err.Error())
	}
	host, err := getEnv("RGE_HOST")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return dbUser{typeDb, user, password, port, name, host}
}

func newMySqlBD(u *dbUser) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", u.User, u.Password, u.Host, u.Port, u.NameDB)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	fmt.Println("conectado a MySql")
}

// newPostgresDB
func newPostgresDB(u *dbUser) {
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", u.User, u.Password, u.Host, u.Port, u.NameDB)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	fmt.Println("conectado a postgres")
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
