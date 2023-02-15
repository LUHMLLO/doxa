package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	Fecha          string
	DeviceName     string
	TempSuperior   float64
	TempIntermedia float64
	TempInferior   float64
	Usuario        string
}

func listAllDevices(db *sql.DB) {

	res, err := db.Query("SELECT * FROM TempReg")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	for res.Next() {
		var fridge Device

		err := res.Scan(&fridge.Fecha, &fridge.DeviceName, &fridge.TempSuperior, &fridge.TempIntermedia, &fridge.TempInferior, &fridge.Usuario)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Fridge:")
		fmt.Println("- name    : ", fridge.DeviceName)
		fmt.Println("- owner   : ", fridge.Usuario)
		fmt.Println("- tempSup : ", fridge.TempSuperior)
		fmt.Println("- tempMid : ", fridge.TempIntermedia)
		fmt.Println("- tempSub : ", fridge.TempInferior)
		fmt.Println("- created : ", fridge.Fecha)
		fmt.Println("")

	}

}

func addNewDevice(db *sql.DB) {

	insert, err := db.Query("INSERT INTO TempReg (`DeviceName`, `TempSuperior`, `TempIntermedia`, `TempInferior`, `Usuario`) VALUES('DevClient - Smart Fridge #03', 9, 6, 3, 'goDev')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}

func main() {
	// D.S.N (Data Source Name) format
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	var (
		host     = "142.93.207.120"
		database = "Smart_Nevera"
		user     = "Nevera"
		password = "0144250809JDR@f"
		dbPath   = fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	)

	db, err := sql.Open("mysql", dbPath)
	if err != nil {
		fmt.Println("error validating sql.Open arguments, dbPath might be incorrect")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping()")
		panic(err.Error())
	}
	fmt.Println("Successful connection to database \n ")

	// call functions here
	//addNewDevice(db)
	listAllDevices(db)

}
