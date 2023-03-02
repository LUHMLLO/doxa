## Tecnica
Leer, aprender , interpretar cualquier programa ya creado que logre el objetivo a implementar,
hacer y re-hacerlo a mi propio modo hasta entender como funciona todo, porque y como mejorarlo
y siempre simplificar. Menos es mas.

## Crea la db en docker

```ps
docker run --name DoxaPostgres -e POSTGRES_DB=doxatec -e POSTGRES_USER=doxadmin -e POSTGRES_PASSWORD=d@x@dm1n -p 5432:5432 -d postgres:latest
```
## API para [ Usuarios || Dispositivos ]:

- CRUD completo basado en ID's donde aplican

- Obtener 1 usuario con los uno o mas parametros/columnas [ Username, Email, Phone ]

- Registra usuarios con los parametros/columnas unicos [ Username, Email, phone ]

- Almacena la password en un hash

- Logea usuarios con los parametros/columnas [ Username, Password ]

- Compara la password ingresada con el hash de la db

- Genera JWT

## Rutas

```go

POST    "/auth/signup"
POST    "/auth/signin"

GET     "/api/users"
POST    "/api/users/create"
GET     "/api/users/read/{id}"
UPDATE  "/api/users/update/{id}"
DELETE  "/api/users/delete/{id}"

GET     "/api/devices"
POST    "/api/devices/create"
GET     "/api/devices/read/{id}"
UPDATE  "/api/devices/update/{id}"
DELETE  "/api/devices/delete/{id}"

```

## Typesafety de la app y la db

```go

// Crea la tabla de usuarios con los parametros a columnas y json
type User struct {
	ID       uuid.UUID  `json:"id"           sql:id varchar(250) primary key`
	Username string     `json:"username"     sql:username varchar(250)`
	Password string     `json:"password"     sql:password varchar(250)`
	Avatar   string     `json:"avatar"       sql:avatar varchar(250)`
	Name     string     `json:"name"         sql:name varchar(250)`
	Email    string     `json:"email"        sql:email varchar(250)`
	Phone    string     `json:"phone"        sql:tempmid decimal`
	Role     string     `json:"role"         sql:role varchar(250):`
	Created  time.Time  `json:"created"      sql:created timestamp`
	Modified time.Time  `json:"modified"     sql:modified timestamp`
}

// Crea la tabla de dispositivos con los parametros a columnas y json
type Device struct {
	ID       uuid.UUID  `json:"id"           sql:id varchar(250) primary key`
	Owner    uuid.UUID  `json:"owner"        sql:owner varchar(250)`
	Name     string     `json:"name"         sql:name varchar(250)`
	TempSup  float64    `json:"temp_sup"     sql:tempsup decimal`
	TempMid  float64    `json:"temp_mid"     sql:tempmid decimal`
	TempSub  float64    `json:"temp_sub"     sql:tempsub decimal`
	Created  time.Time  `json:"created"      sql:created timestamp`
	Modified time.Time  `json:"modified"     sql:modified timestamp`
}

```