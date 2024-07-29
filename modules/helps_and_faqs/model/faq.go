package modelFAQ

import (
	"fastFood/common"
	"strings"
)

const (
	EntityName = "FAQ"
)

var (
	ErrRequestIsBlank = "request cannot be blank"
	ErrAnswerIsBlank  = "answer cannot be blank"
)

type FAQ struct {
	common.SQLModel
	Question string `json:"question" gorm:"column:question;"`
	Answer   string `json:"answer" gorm:"column:answer;"`
}

func (FAQ) TableName() string {
	return "helps_and_faqs"
}

type FAQCreate struct {
	Question string `json:"question" gorm:"column:question;"`
	Answer   string `json:"answer" gorm:"column:answer;"`
}

func (FAQCreate) TableName() string {
	return FAQ{}.TableName()
}

func (faq *FAQCreate) Validate() error {
	faq.Question = strings.TrimSpace(faq.Question)
	faq.Answer = strings.TrimSpace(faq.Answer)

	if faq.Question == "" {
		return ErrValidateRequest(ErrRequestIsBlank, "ERR_REQUEST_IS_BLANK")
	}

	if faq.Answer == "" {
		return ErrValidateRequest(ErrAnswerIsBlank, "ERR_ANSWER_IS_BLANK")
	}

	return nil
}
