package controllers

import (
	// "errors"
	"github.com/gin-gonic/gin"
	// "gopkg.in/mgo.v2/bson"
	"net/http"
	"netlui-go-server/conn"
	payloads "netlui-go-server/models/Student"

	//
	// "fmt"
	"math/rand"
	// "strconv"
	"time"
)

const UserCollection = "users"

func generateId(min, max int) int { // check wether concurrency is required here
	return min + rand.Intn(max-min)
}

func generateUniqId() int {
	rand.Seed(time.Now().UnixNano())
	uid := generateId(0, 2000)

	return uid
}

func Authorize(c *gin.Context) {
	db := conn.GetMongoDB()
	register := payloads.RegisterModel{}
	err := c.Bind(&register)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "failed to create an account"})
		return
	}

	register.UserId = generateUniqId()
	error := db.C(UserCollection).Insert(register)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed", "user": &register})
	}

	c.JSON(http.StatusOK, gin.H{"status": "authorization sucess!", "message": "your account was created"})

}
