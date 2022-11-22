package main

import (
	"fmt"
	"log"
	"net/http"
	"onestep/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Routes(r)
	http.Handle("/", r)

	if err := http.ListenAndServe(":9010", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Print("listening on port 9010...")
}
