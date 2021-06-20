# Restaurante-Gorm-Echo

## Requisitos

Para el correcto funcionamiento del programa se deberá contar con un sistema operativo Windows 7 o superior y tener el motor de base de datos MySql o PostgreSql

## Estructura de archivos

Para usar el programa se deberá crear una carpeta llamada cmd y otra llamada public, ambas deben estar en el mismo directorio. El programa usa rutas relativas, por lo que se puede seleccionar cualquier directorio.

A continuación se muestra un ejemplo de la estructura de archivos.

![Estructura de archivos](Estructura_archivos.JPG "Estructura de archivos")

### Carpeta cmd

Dentro de la carpeta cmd se encontrará el ejecutable o archivo main.go, además de otra carpeta llamada certificates

#### Carpeta certificates

Aquí se almacenarán todas las credenciales necesarias para el correcto funcionamiento del programa, los archivos necesarios son:

- app.rsa

- app.rsa.pub

- cookieKey

- db.json

- email.json

### Carpeta public

Dentro se almacenará la carpeta views, en la cual estará todo el código html, js y css.

La carpeta template contendrá el archivo confirm.html, el cual será usado como plantilla al momento de enviar los códigos de confirmación.

También contendrá la carpeta views/assets/img/products/,  en dicha carpeta se almacenarán las imágenes de los productos de forma automática.

## Credenciales

En total se debe de contar con 5 archivos, los cuales se mostraran a continuación, las claves rsa pueden ser generadas en <https://travistidwell.com/jsencrypt/demo/>

### app.rsa

Este archivo contendrá la clave rsa privada

### app.rsa.pub

Este archivo contendrá la clave rsa pública

### cookieKey

Este archivo contendrá una clave con una longitud de 32 caracteres

### db.json

Este archivo debe tener la siguiente estructura:

```json
{
  "type_db":"",
  "user":"",
  "password":"",
  "port":"",
  "name_db":""
}
```

- **type_db:** especifica el tipo de base de datos, puede ser *POSTGRES* o *MYSQL*, seleccionar POSTGRES realizará una conexión con la base de datos en PostgreSql, en cambio si se selecciona MYSQL se realizará la conexión con MySql.

- **user:** selecciona el usuario con el cual se accedera a la base de datos.

- **password:** selecciona la contraseña del usuario.

- **port:** selecciona el puerto con el cual establece la conexión a la base de datos.

- **name_db:** selecciona el nombre de la base de datos a la cual se conectara.

### email.json

Este archivo debe tener la siguiente estructura:

```json
{
  "email":"",
  "password":""
}
```

- **email:** selecciona el email del cual seran enviados los codigos de verificacion.

- **password:** selecciona la contraseña del email.

## Migraciones

El programa realiza las migraciones de todos los modelos dentro de la función AutoMigrate, si se desean agregar en un futuro nuevos modelos y que estos se almacenen en la base de datos será necesario agregar dicho modelo a la función AutoMigrate que se encuentra en el archivo main.go

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

## Contraseñas

Las contraseñas de usuarios deben de contener al menos una mayúscula, una minúscula y un carácter especial, además de tener una longitud mayor o igual a 8

## Primera ejecución del programa

Cuando se ejecuta el programa por primera vez se realizarán las migraciones de las tablas de forma automática, además de que se solicitará un correo electrónico y una contraseña para crear al usuario con rol de Dueño, el cual no tiene restricciones sobre el uso del programa. Posteriormente se preguntará si se quieren utilizar seeders para poblar las bases de datos y ver el funcionamiento del programa.
Mientras no se haya confirmado el usuario con rol de Dueño el programa pedirá que se ingresen las credenciales de dicho usuario y enviará el código de verificación.

### Seeders

El programa puede ejecutar seeders para poblar la tabla de productos, direcciones, establecimientos, usuarios y pedidos, se pedirá al usuario que ingrese la cantidad de elementos en cada tabla y se llenarán con datos aleatorios de un faker.

## Permisos

Cada permiso está identificado por un id y posee una descripción sobre su funcionamiento

1. Sin restricciones

2. CRUD estaclimientos

3. Asignar roles de menor jerarquía

4. CRUD productos

5. Tomar orden en un establecimiento

6. Confirmar Pago

7. Ver todas las órdenes sin completar del local

8. CRUD métodos de pago

9. Ver facturas

10. Asignar empleado a un establecimiento

11. Mostrar todos los empleados del establecimiento

12. Mostrar todos los empleados

13. CRUD Roles

14. CRUD Tables

15. Mostrar ordenes pendientes del empleado

16. Asignar rol de menor jerarquía de empleado dentro de un establecimiento

## Roles

Los roles son un conjunto de permisos que puede tener un usuario. A continuación se muestran los roles creados hasta el momento y sus respectivos permisos mediante la ID de cada permiso

- **Dueño**
  - 1
- **Admin**
  - 2
  - 3
  - 4
  - 8
  - 11
  - 12
  - 13
- **Manager**
  - 9
  - 10
  - 11
  - 16
- **Mesero**
  - 5
  - 6
  - 15
- **Cocina**
  - 7
