package server

import (
	"log"
	"net/http"

	"github.com/katungi/gut/internal/bundler"
)

func Start() {
	bundleErr := bundler.Bundle("../../example/src/index.js", "./public/out.js")
	if bundleErr != nil {
		log.Fatalf("could not bundle the JavaScript")
	}
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	log.Println("Listening to on port :6699")

	err := http.ListenAndServe(":6699", nil)

	if err != nil {
		log.Fatal(err)
	}
}
