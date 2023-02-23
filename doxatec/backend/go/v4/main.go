package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewApiServer(":3000", store)
	server.Start()
}

/*
	Hey vlad, klk, sé que no he dado mucho seguimiento,
	este finde mi cerebro no daba pa' nada so durante el fin de semana no avance mucho,
	desde el lunes he estado testeando diversas formas de manejar el usuario de manera simplificada para nuestro objetivo,
	al mismo tiempo eh hecho como 3 backends diferentes en go xD
	pero si no era que funcionaban estrictamente por api, simplemente funcionaba de manera estatica,
	de igual forma, entre prueba y error creo que ya eh conseguido una mejor manera
	que vendria siendo asi:

	-->
*/
