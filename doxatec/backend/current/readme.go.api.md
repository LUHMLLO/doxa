# Doxapi

## Routes

```js
HOST:  `http://142.93.207.120:3000`

GET:    `/health`
POST:   `/api/auth/signup`
POST:   `/api/auth/signin`
GET:    `/api/auth/signature`
GET:    `/api/auth/mydevices`
POST:   `/api/auth/signout`

GET:    `/api/master/users/all`
POST:   `/api/master/users/insert`
GET:    `/api/master/users/read/${id}`
UPDATE: `/api/master/users/update/${id}`
DELETE: `/api/master/users/delete/${id}`

GET:    `/api/master/devices/all`
POST:   `/api/master/devices/insert`
GET:    `/api/master/devices/read/${id}`
UPDATE: `/api/master/devices/update/${id}`
DELETE: `/api/master/devices/delete/${id}`

```

## Typesafety

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

type Device struct {
 ID       uuid.UUID  `json:"id"         sql:id varchar(250) primary key`
 Owner    uuid.UUID  `json:"owner"      sql:owner varchar(250)`
 Name     string     `json:"name"       sql:name varchar(250)`
 PIN      string     `json:"pin"        sql:pin varchar(250)`
 TempSup  float64    `json:"temp_sup"   sql:tempsup decimal`
 TempMid  float64    `json:"temp_mid"   sql:tempmid decimal`
 TempSub  float64    `json:"temp_sub"   sql:tempsub decimal`
 Created  time.Time  `json:"created"    sql:created timestamp`
 Modified time.Time  `json:"modified"   sql:modified timestamp`
}

```
