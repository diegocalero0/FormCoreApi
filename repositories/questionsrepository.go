package repositories

import (
	"formcoreapi/api"
	"formcoreapi/dtos"
)

type QuestionsRepository struct{}

// function that allows to fetch the questions from Teamcore Api
func (r *QuestionsRepository) GetQuestions() (*dtos.GetQuestionResponseDTO, error) {
	return api.Get[dtos.GetQuestionResponseDTO]()
}
