# Doxapi

## Dependencias

- ### Base de datos Postgres

    ```docker
    docker run --name DoxaDB -e POSTGRES_DB=doxatec -e POSTGRES_USER=doxadmin -e POSTGRES_PASSWORD=d@x@dm1n -p 5432:5432 -d postgres:latest
    ```

## Rutas

- ### Host IP Address

    ```js
    HOST:  `http://142.93.207.120:3000`
    ```

- ### Host Domain

    ```js
    HOST:  `<http://142.93.207.120:3000>`
    ```

- ### Host endpoints

    ```js
    GET:    `init/health`
    ```

    ```js
    POST:   `/api/user/signup`
    POST:   `/api/user/signin`
    GET:    `/api/user/signature`
    POST:   `/api/user/signout`
    
    GET:    `/api/user/mydevices`
    ```

    ```js
    GET:    `/api/master/clients/list`
    GET:    `/api/master/clients/list/all`
    GET:    `/api/master/clients/list/:column`

    POST:   `/api/master/clients/create`
    GET:    `/api/master/clients/read/:id`
    UPDATE: `/api/master/clients/update/:id`
    DELETE: `/api/master/clients/delete/:id`
    ```

    ```js
    GET:    `/api/master/users/list`
    GET:    `/api/master/users/list/all`
    GET:    `/api/master/users/list/:column`

    POST:   `/api/master/users/create`
    GET:    `/api/master/users/read/:id`
    UPDATE: `/api/master/users/update/:id`
    DELETE: `/api/master/users/delete/:id`
    ```

    ```js
    GET:    `/api/master/devices/list`
    GET:    `/api/master/devices/list/all`
    GET:    `/api/master/devices/list/:column`

    POST:   `/api/master/devices/create`
    GET:    `/api/master/devices/read/:id`
    UPDATE: `/api/master/devices/update/:id`
    DELETE: `/api/master/devices/delete/:id`
    ```

## Typesafety

- ### Clients

    ```go
    type Clients struct {
        ID       uuid.UUID  `json:"id"`         `sql:id varchar(250) primary key`
        Name     string     `json:"name"`       `sql:name varchar(250)`
        Email    string     `json:"email"`      `sql:email varchar(250)`
        Phone    string     `json:"phone"`      `sql:phone decimal`
        Created  time.Time  `json:"created"`    `sql:created timestamp NULL`
        Modified time.Time  `json:"modified"`   `sql:modified timestamp NULL`
    }
    ```

- ### Users

    ```go
    type Users struct {
        ID       uuid.UUID  `json:"id"`         `sql:id varchar(250) primary key`
        Client   Clients    `json:"client"`     `sql:client varchar(250)`
        Avatar   string     `json:"avatar"`     `sql:avatar varchar(250)`
        Username string     `json:"username"`   `sql:username varchar(250)`
        Password string     `json:"password"`   `sql:password varchar(250)`
        Role     string     `json:"role"`       `sql:role varchar(250)`
        Created  time.Time  `json:"created"`    `sql:created timestamp NULL`
        Modified time.Time  `json:"modified"`   `sql:modified timestamp NULL`
        Accessed time.Time  `json:"accessed"`   `sql:accessed timestamp NULL`
    }
    ```

- ### Devices

    ```go
    type Devices struct {
        ID       uuid.UUID  `json:"id"`         `sql:id varchar(250) primary key`
        Client   Clients    `json:"client"`     `sql:client varchar(250)`
        Name     string     `json:"name"`       `sql:name varchar(250)`
        PIN      string     `json:"pin"`        `sql:pin varchar(250)`
        TempSup  float64    `json:"temp_sup"`   `sql:temp_sup numeric(4,2)`
        TempMid  float64    `json:"temp_mid"`   `sql:temp_mid numeric(4,2)`
        TempSub  float64    `json:"temp_sub"`   `sql:temp_sub numeric(4,2)`
        Created  time.Time  `json:"created"`    `sql:created timestamp NULL`
        Modified time.Time  `json:"modified"`   `sql:modified timestamp NULL`
    }

    ```

- ### Plans

    ```go

    type User struct {
    ID       uuid.UUID  `json:"id"         sql:id varchar(250) primary key`
    Username string     `json:"username"   sql:username varchar(250)`
    Password string     `json:"password"   sql:password varchar(250)`
    Avatar   string     `json:"avatar"     sql:avatar varchar(250)`
    Name     string     `json:"name"       sql:name varchar(250)`
    Email    string     `json:"email"      sql:email varchar(250)`
    Phone    string     `json:"phone"      sql:tempmid decimal`
    Role     string     `json:"role"       sql:role varchar(250):`
    Created  time.Time  `json:"created"    sql:created timestamp`
    Modified time.Time  `json:"modified"   sql:modified timestamp`
    }
    ```

- ### Payments

    ```go

    type User struct {
    ID       uuid.UUID  `json:"id"         sql:id varchar(250) primary key`
    Username string     `json:"username"   sql:username varchar(250)`
    Password string     `json:"password"   sql:password varchar(250)`
    Avatar   string     `json:"avatar"     sql:avatar varchar(250)`
    Name     string     `json:"name"       sql:name varchar(250)`
    Email    string     `json:"email"      sql:email varchar(250)`
    Phone    string     `json:"phone"      sql:tempmid decimal`
    Role     string     `json:"role"       sql:role varchar(250):`
    Created  time.Time  `json:"created"    sql:created timestamp`
    Modified time.Time  `json:"modified"   sql:modified timestamp`
    }
    ```
