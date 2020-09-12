package quiz

// refModel "netlui-go-server/models"

type QuizResponsePayload struct {
	QuizId      int    `json:"quiz_id"`
	ChapterName string `json:"chapter_name"`
	SubjectName string `json:"subject_name"`
	// nested struct
	QnAs []QnA
}
