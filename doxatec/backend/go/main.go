package main

import (
	"MyGoSql/api"
	"MyGoSql/storage"
	"flag"
	"fmt"
	"log"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()

	listenAddress := flag.String("listenAddress", ":5000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddress, store)
	fmt.Println("Storage: ", reflect.TypeOf(*store))
	fmt.Println("Server running on port: ", *listenAddress)
	fmt.Printf("Http Address: http://localhost%s/devices\n", *listenAddress)
	fmt.Printf("API Endpoint: http://localhost%s/owners\n", *listenAddress)
	fmt.Printf("API Endpoint: http://localhost%s/users\n", *listenAddress)
	fmt.Println("booting time: ", time.Since(startTime))
	log.Fatal(server.Start())
}
