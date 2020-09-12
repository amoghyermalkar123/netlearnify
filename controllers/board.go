package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"netlui-go-server/conn"

	"gopkg.in/mgo.v2/bson"

	dbmodels "netlui-go-server/models/domain"
	// dbmodels "netlui-go-server/models/quiz"
	//  "fmt"
	// "math/rand"
	// "strconv"
	// "time"
)

const PinboardCollection = "pinboard"
const SubjectProgressCollection = "subjectAnalysis"

func GetBoardData(c *gin.Context) {
	db := conn.GetMongoDB()
	// instance declarations
	domainIns := dbmodels.DomainAnalysis{}
	// subjectProgressIns := dbmodels.SubjectWiseProgress{}
	requestPayload := dbmodels.BoardRequestPayload{}
	pinboard := dbmodels.Board{}

	subjectDocIdList := make([]int, 0)

	err := c.Bind(&requestPayload)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "payload problem"})
	}

	// retrieve board data
	findQuery := bson.M{"userid": requestPayload.StudentID}
	// concurrently retrieve quiz data
	findOp := db.C(PinboardCollection).Find(findQuery).One(&domainIns)

	if findOp != nil {
		c.JSON(http.StatusOK, gin.H{"status": "database problem"})

	} else {
		pinboard.UserId = requestPayload.StudentID
	}

	subjectDocIdList = append(subjectDocIdList, domainIns.Subjects...)

	for i := 0; i < len(subjectDocIdList); i++ {
		// make a function call
		subjectAnalysisDOC := searchSubjectDocuments(subjectDocIdList[i])
		// append returned doc object into response payloads array of struct fild
		pinboard.Subjects = append(pinboard.Subjects, subjectAnalysisDOC)

	}

	// getRecentQuizData := bson.M{"": -1}
	quizRetErr := db.C(StudentQuizRecords).Find(bson.M{}).Sort("-_id").Limit(4).All(&pinboard.Quizes)

	if quizRetErr != nil {
		c.JSON(http.StatusOK, gin.H{"status": "test fail", "board": "fail to load quizes"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "test success", "board": &pinboard})

	}

}

func searchSubjectDocuments(docID int) (doc dbmodels.SubjectWiseProgress) {
	db := conn.GetMongoDB()
	subjectDocINS := dbmodels.SubjectWiseProgress{}
	findDocQuery := bson.M{"subjectDocid": docID}
	err := db.C(SubjectProgressCollection).Find(findDocQuery).One(&subjectDocINS)
	if err != nil {
		return subjectDocINS
	}
	return subjectDocINS
}

// func UpdateBoard() {

// }

// func ProgressCalculator([]int) (rate int) {

// }

// func GetQuizData(id int) {

// }

/*
find().sort({"_id": -1}).limit(4).pretty()

*/
