package main

import (
	"cribl-logger/internal/controllers"
	"fmt"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
)

var defaultPort string = ":8080"

func main() {
	Run()
}
func Run() {
	ctlr := controllers.NewMainController()

	// Setup routing
	r := chi.NewRouter()

	r.Get("/logs", ctlr.Logs)

	port := getPort()
	fmt.Println("Listening on port: ", port)
	http.ListenAndServe(port, r)
}

func getPort() string {
	port := os.Getenv("CRIBL_PORT")
	if port == "" {
		return defaultPort
	} else {
		return port
	}
}
