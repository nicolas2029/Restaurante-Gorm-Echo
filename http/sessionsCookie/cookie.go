package sessionsCookie

import (
	"os"
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

func NewCookieStore() error {
	var err error
	once.Do(func() {
		loadKey()
		cookie = gormstore.New(storage.DB(), key)
		quit := make(chan struct{})
		go cookie.PeriodicCleanup(1*time.Hour, quit)
	})
	return err
}

func loadKey() {
	keys, _ := os.LookupEnv("RGE_COOKIE_KEY")
	key = []byte(keys)
}
