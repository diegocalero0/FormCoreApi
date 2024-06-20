package main

import (
	"fmt"
	"formcore/api/repositories"
	"log"

	"github.com/joho/godotenv"
)

func initializeDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
}

func main() {
	initializeDotenv()

	repository := repositories.QuestionsRepository{}
	questions, err := repository.GetQuestions()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(questions.Date)
	}

}
