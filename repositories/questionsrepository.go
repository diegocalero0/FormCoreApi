package repositories

import (
	"formcore/api/api"
	"formcore/api/dtos"
)

type QuestionsRepository struct{}

// function that allows to fetch the questions from Teamcore Api
func (r *QuestionsRepository) GetQuestions() (*dtos.GetQuestionResponseDTO, error) {
	return api.Get[dtos.GetQuestionResponseDTO]()
}
