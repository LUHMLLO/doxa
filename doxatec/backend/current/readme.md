# Doxapi

## Dependencias

- ### Base de datos Postgres

    [postgres docker](./readme.postgres.sh)

## Rutas

- ### Host endpoints

    [api endpoints](./readme.endpoints.json)

- ### Host IP Address

    ```js
    HOST:  `http://142.93.207.120:3000`
    ```

- ### Host Domain

    ```js
    HOST:  `<http://142.93.207.120:3000>`
    ```

## Typesafety

- ### Clients

    ```go
    type Clients struct {
        ID       uuid.UUID  `json:"id"`
        Name     string     `json:"name"`
        Email    string     `json:"email"`
        Phone    string     `json:"phone"`
        Created  time.Time  `json:"created"`
        Modified time.Time  `json:"modified"`
    }
    ```

- ### Users

    ```go
    type Users struct {
        ID       uuid.UUID  `json:"id"`
        Client   Clients    `json:"client"`
        Avatar   string     `json:"avatar"`
        Username string     `json:"username"`
        Password string     `json:"password"`
        Role     string     `json:"role"`
        Created  time.Time  `json:"created"`
        Modified time.Time  `json:"modified"`
        Accessed time.Time  `json:"accessed"`
    }
    ```

- ### Devices

    ```go
    type Devices struct {
        ID       uuid.UUID      `json:"id"`
        Client   Clients        `json:"client"`
        Name     string         `json:"name"`
        PIN      string         `json:"pin"`
        Temps    Temperatures  `json:"temps"`
        Created  time.Time      `json:"created"`
        Modified time.Time      `json:"modified"`
    }

    ```

- ### Temperatures

    ```go
    type Temperatures struct {
        ID       uuid.UUID  `json:"id"`
        Device   Devices    `json:"device"`
        TempSup  float64    `json:"temp_sup"`
        TempMid  float64    `json:"temp_mid"`
        TempSub  float64    `json:"temp_sub"`
        Created  time.Time  `json:"created"`
    }

    ```
