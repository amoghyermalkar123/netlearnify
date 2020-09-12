package domain

import (
	model "netlui-go-server/models/quiz"
)

type Board struct {
	UserId   int                       `json:"userid"`
	Subjects []SubjectWiseProgress     `json:"subjects"`
	Quizes   []model.StudentQuizRecord `json:"quizes"`
}
