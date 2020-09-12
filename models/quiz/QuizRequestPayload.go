package quiz

// import (
//   refModel "netlui-go-server/models"
//
// )

type QuizRequestPayload struct {
	UserId      int    `json:"user_id"`
	ChapterName string `json:"chapter_name"`
	SubjectName string `json:"subject_name"`
}
