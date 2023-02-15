package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "Nevera"
	password string = "0144250809JDR@f"
	hostname string = "142.93.207.120"
	dbname   string = "Smart_Nevera"
)

var (
	DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
)

type Device struct {
	Name    string
	Owner   string
	TempSup float64
	TempMid float64
	TempSub float64
	Created string
}

func CheckConnection() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		panic(err.Error())
	}
	defer db.Close()
}

func CreateDatabase(db *sql.DB) {
	res, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		fmt.Printf("Error %s when creating DB\n", err.Error())
		return
	}

	number, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error %s when fetching rows\n", err.Error())
		return
	}

	fmt.Printf("rows affected %d\n", number)
	db.Close()
}

func Connect(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping()")
		panic(err.Error())
	}
	fmt.Println("Successful connection to database \n ")
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Devices (Name VARCHAR(250), Owner VARCHAR(250), TempSup FLOAT, TempMid FLOAT, TempSub FLOAT, Created TIMESTAMP)")
	if err != nil {
		panic(err)
	}
}

func InsertToDevices(db *sql.DB) {
	_, err := db.Exec("INSERT INTO Devices (`Name`,`Owner`,`TempSup`, `TempMid`, `TempSub`, `Created`) VALUES('Nevera 0003','DevClient', 13.5, 8.6, 2.1, NOW())")
	if err != nil {
		panic(err)
	}
}

func ListAllDevices(db *sql.DB) {

	res, err := db.Query("SELECT * FROM Devices")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	for res.Next() {
		var column Device

		err := res.Scan(&column.Name, &column.Owner, &column.TempSup, &column.TempMid, &column.TempSub, &column.Created)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Device")
		fmt.Println("- name    : ", column.Name)
		fmt.Println("- owner   : ", column.Owner)
		fmt.Println("- tempSup : ", column.TempSup)
		fmt.Println("- tempMid : ", column.TempMid)
		fmt.Println("- tempSub : ", column.TempSub)
		fmt.Println("- created : ", column.Created)
		fmt.Println("")
	}

}

func main() {
	CheckConnection()

	db, _ := sql.Open("mysql", DSN)
	defer CreateDatabase(db)

	Connect(db)
	CreateTable(db)
	//InsertToDevices(db)
	ListAllDevices(db)
}
