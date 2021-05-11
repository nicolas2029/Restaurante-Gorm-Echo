package sessionsCookie

import (
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/wader/gormstore/v2"
)

var (
	once   sync.Once
	cookie *gormstore.Store
	key    []byte
)

func Cookie() *gormstore.Store {
	return cookie
}

func NewCookieStore(file string) error {
	var err error
	once.Do(func() {
		err = loadKey(file)
		if err != nil {
			log.Fatalf("No se pudo cargar la clave de cookies %v", err)
		}
		cookie = gormstore.New(storage.DB(), key)
		quit := make(chan struct{})
		go cookie.PeriodicCleanup(1*time.Hour, quit)
	})
	return err
}

func loadKey(file string) error {
	var err error
	key, err = ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return nil
}
