package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gbodra/swimming-calculator-api/controller"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type App struct {
	Router *mux.Router
	Port   string
}

func (a *App) Initialize() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env")
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// Management routes
	a.Router.HandleFunc("/health", controller.HealthCheck).Methods("GET")

	// Race routes
	a.Router.HandleFunc("/race", controller.RacePost).Methods("POST")
}

func (a *App) Run() {
	port := getPort()
	log.Println("App running on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Println("$PORT not set. Falling back to default " + port)
	}

	return port
}
