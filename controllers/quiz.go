package controllers

import (
	// "errors"
	"net/http"
	"netlui-go-server/conn"
	payloads "netlui-go-server/models/quiz"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	// "math/rand"
	// "strconv"
	// "time"
)

var QuizCollection = "Quizes"
var QnACollection = "questions_answers"
var StudentQuizRecords = "quiz_records"

func StoreQnA(c *gin.Context) {
	db := conn.GetMongoDB()
	qnaInstance := payloads.QnA{}
	err := c.Bind(&qnaInstance)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "payload problem"})
		return
	}

	qnaInstance.QnA_Id = generateUniqId()
	error := db.C(QnACollection).Insert(qnaInstance)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed", "message": "failed to store the document"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "sucess!", "message": "your doc has been stored"})

}

func StoreStudentQuiz(c *gin.Context) {
	// basic declarations
	db := conn.GetMongoDB()
	studentQuizRecord := payloads.StudentQuizRecord{}

	// binding
	bindErr := c.Bind(&studentQuizRecord)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "There was some problem with the payload"})
		return
	}

	// persist the payload directly
	insertOpError := db.C(StudentQuizRecords).Insert(studentQuizRecord)

	// database error checking
	if insertOpError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "There was some problem submitting your quiz"})
		return
	}

	// continuing without any errors ...
	// sending success response
	c.JSON(http.StatusOK, gin.H{"response": &studentQuizRecord, "message": "Stored Quiz Succesfuly!!"})

}

func searchQnADoc(id int) (qnaIns payloads.QnA) {
	db := conn.GetMongoDB()
	quizQNAINS := payloads.QnA{}
	findDocQuery := bson.M{"qna_id": id}
	err := db.C(QnACollection).Find(findDocQuery).One(&quizQNAINS)
	if err != nil {
		return quizQNAINS
	}
	return quizQNAINS

}

func LoadQuiz(c *gin.Context) {
	// TODO : ADD A FILTER FUNCTIONALITY TO THIS HANDLER FUNCTION
	db := conn.GetMongoDB()
	// basic instance declarations
	// userQuizRequestInstance := payloads.QuizRequestPayload{}
	responsePayloadIns := payloads.QuizResponsePayload{}
	quizIns := payloads.QuizModel{}

	// bindErr := c.Bind(&userQuizRequestInstance)

	// if bindErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Response": "Problem wityour payload"})
	// }
	// a temporary array that contains all QnA Doc IDs of the retrieved Quiz Document
	qnaIDs := make([]int, 0)
	// retrieve random QuizModel
	findQuizDocQuery := bson.M{"quiz_id": 1}
	err := db.C(QuizCollection).Find(findQuizDocQuery).One(&quizIns) //

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Response": "Failed to Load Quiz"})
	} else {
		// load that data in the response payload
		// intID, _ := strconv.Atoi()
		responsePayloadIns.QuizId = quizIns.QuizId
		responsePayloadIns.ChapterName = quizIns.ChapterName
		responsePayloadIns.SubjectName = quizIns.SubjectName
	}

	qnaIDs = append(qnaIDs, quizIns.QnA_Ids...)
	// retrieve QnA ocuments from the id array of Quiz Model that was retrieved
	for i := 0; i < len(qnaIDs); i++ {
		// make a function call
		questionAnswerDOC := searchQnADoc(qnaIDs[i])
		// append returned doc object into response payloads array of struct fild
		responsePayloadIns.QnAs = append(responsePayloadIns.QnAs, questionAnswerDOC)

	}
	// load those in the array of response payload

	// success response
	c.JSON(http.StatusOK, gin.H{"status": "Your Quiz is Ready", "Quiz": &responsePayloadIns})

}
