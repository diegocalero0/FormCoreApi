package controllers

import (
	"encoding/json"
	"formcore/api/models"
	"formcore/api/services"
	"net/http"
)

type QuestionsControler struct {
	QuestionsService services.QuestionsService
}

func (c *QuestionsControler) Get(w http.ResponseWriter, r *http.Request) {
	questions, err := c.QuestionsService.GetQuestions()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errorResponse := models.ErrorModel{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(questions)
	}
}
