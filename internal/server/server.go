package server

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/katungi/gut/internal/bundler"
)

func Start() {
	entryFile := filepath.Join("example", "src", "index.js")
	outputFile := filepath.Join("public", "out.js")
	bundleErr := bundler.Bundle(entryFile, outputFile)
	if bundleErr != nil {
		log.Printf(
			"could not bundle the JavaScript files: %v\n",
			bundleErr,
		)
	}
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	log.Println("Listening to on port :6699")

	err := http.ListenAndServe(":6699", nil)

	if err != nil {
		log.Fatal(err)
	}
}
