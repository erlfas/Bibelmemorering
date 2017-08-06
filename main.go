package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web server ...")

	http.Handle("/", &indexHandler{filename: "index.html"})
	http.Handle("/index", &indexHandler{filename: "index.html"})
	http.Handle("/result", &resultHandler{filename: "result.html"})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	log.Println("Stopping ...")
}
