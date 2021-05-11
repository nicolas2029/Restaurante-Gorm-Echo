# Restaurante-Gorm-Echo

## Conexion a la base de datos

```go
    storage.New(storage.Postgres)
```

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

## Password

Las contraseñas deben de contener al menos una mayuscula, una minuscula y un caracter especial, ademas de tener una longitud mayor o igual a 8

## Establishment

Cada permiso esta identificado por un id y posee una descripcion sobre su funcionamiento

1. Sin restricciones

2. CRUD estaclimientos

3. Asignar roles de menor jerarquia

4. CRUD productos

5. Tomar orden en un establecimiento

6. Confirmar Pago

7. Ver todas las ordenes sin completar del local

8. CRUD metodos de pago

9. Ver facturas

10. Asignar empleado a un establecimiento

11. Mostrar todos los empleados del establecimiento

12. Mostrar todos los empleados

13. CRUD Roles

14. CRUD Tables

## Roles

- **owner**
  - 1
- **admin**
  - 2
  - 3
  - 4
  - 8
  - 11
  - 12
  - 13
- **manager**
  - 3
  - 9
  - 10
  - 11
- **waiter**
  - 5
  - 6
  - 7
- **chef**
  - 7
