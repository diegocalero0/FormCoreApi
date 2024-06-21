package services

import (
	"formcoreapi/models"
	"formcoreapi/repositories"
	"formcoreapi/utils"
	"strconv"
)

type QuestionsService struct {
	QuestionsRepository repositories.QuestionsRepository
}

// Function that fetch the questions from the repository and transforms the response
func (s *QuestionsService) GetQuestions() (*models.GetQuestionResponseModel, error) {
	getQuestionsResponse, err := s.QuestionsRepository.GetQuestions()
	if err != nil {
		return nil, err
	} else {
		Info := make([]models.QuestionModel, len(getQuestionsResponse.Data))
		for index, questionDto := range getQuestionsResponse.Data {

			intQuestionId, err := strconv.Atoi(questionDto.QuestionId)

			if err != nil {
				intQuestionId = -1
			}

			Info[index] = models.QuestionModel{
				QuestionId: intQuestionId,
				Question:   questionDto.Question,
			}
		}

		date, err := utils.TransformDate(getQuestionsResponse.Date)

		if err != nil {
			return nil, err
		}

		return &models.GetQuestionResponseModel{
			Title:      "Preguntas sobre cultura general",
			Date:       date,
			Info:       Info,
			ApiVersion: 1,
		}, nil
	}
}
