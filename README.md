# Restaurante-Gorm-Echo

## Migraciones

```go
    storage.DB().AutoMigrate(
		&model.Product{},
		&model.Address{},
		&model.Permission{},
		&model.Establishment{},
		&model.Rol{},
		&model.User{},
		&model.Table{},
		&model.Pay{},
		&model.Order{},
		&model.OrderProduct{},
	)
```
