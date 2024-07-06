package server

import (
	"log"
	"net/http"
)

func Start() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	log.Println("Listening to on port :6699")

	err := http.ListenAndServe(":6699", nil)

	if err != nil {
		log.Fatal(err)
	}
}
