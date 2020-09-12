package controllers

import (
	// "errors"
	"net/http"
	"netlui-go-server/conn"
	payloads "netlui-go-server/models/Student"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	//
	// "fmt"
	// "math/rand"
	// // "strconv"
	// "time"
)

func Authenticate(c *gin.Context) {
	db := conn.GetMongoDB()
	login := payloads.LoginPayload{}
	err := c.Bind(&login)
	reg_model := payloads.RegisterModel{}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "failed to login"})
		return
	}

	authenticateQuery := bson.M{
		"$and": []bson.M{
			bson.M{"email": &login.Email},
			bson.M{"password": &login.Password},
		},
	}
	error := db.C(UserCollection).Find(authenticateQuery).One(&reg_model)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "cannot find account", "error": &error})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "login succesfull", "userId": &reg_model.UserId})
	}

}
