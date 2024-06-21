package dtos

type GetQuestionResponseDTO struct {
	Date string        `json:"date"`
	Data []QuestionDTO `json:"data"`
}
