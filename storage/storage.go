package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
}

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
// Drivers
const (
	Postgres string = "POSTGRES"
	MySql    string = "MYSQL"
)

// New create the connection with db
func New(file string) {
	once.Do(func() {
		u := loadFileDB(file)
		switch u.TypeDB {
		case Postgres:
			newPostgresDB(&u)
		case MySql:
			newMySqlBD(&u)
		}
	})
}
func loadFileDB(file string) dbUser {
	var err error
	m, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("no se pudo cargar las credenciales de DB: %v", err)
	}
	u := dbUser{}
	err = json.Unmarshal(m, &u)
	if err != nil {
		log.Fatalf("fallo en unmarshal DB: %v", err)
	}
	return u
}

func newMySqlBD(u *dbUser) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", u.User, u.Password, u.Port, u.NameDB)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	fmt.Println("conectado a MySql")
}

func newPostgresDB(u *dbUser) {
	var err error
	//"postgres://admin_restaurant:RestAuraNt_pgsql.561965697@localhost:5433/restaurante_gorm_echo?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", u.User, u.Password, u.Port, u.NameDB)
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
