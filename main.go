package main

import (
	"formcoreapi/controllers"
	"formcoreapi/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Function that loads all enviroment variables from .env file
func initializeDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v"+"Porfavor cree un archivo .env en la raíz del proyecto con las variables TEAMCORE_API_URL y TEAMCORE_API_TOKEN", err)
	}
}

func main() {
	initializeDotenv()
	questionsController := controllers.QuestionsControler{}
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/questions", middlewares.AuthMiddleware(http.HandlerFunc(questionsController.Get)))
	router.HandleFunc("/generateToken", middlewares.GenerateTokenHandler)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
