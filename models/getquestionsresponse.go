package models

type GetQuestionResponseModel struct {
	Title      string          `json:"titulo"`
	Date       string          `json:"dia"`
	Info       []QuestionModel `json:"info"`
	ApiVersion int             `json:"api_version"`
}
