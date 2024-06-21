package dtos

type QuestionDTO struct {
	QuestionId string      `json:"question_id"`
	Question   string      `json:"question"`
	Answers    []AnswerDTO `json:"answers"`
}
