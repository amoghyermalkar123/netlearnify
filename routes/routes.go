package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// import controllers/ hadnler functions
	handler "netlui-go-server/controllers"
)

func StartGin() {
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		// AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	api := router.Group("/apiv1")
	{
		api.POST("/board/", handler.GetBoardData)
		api.POST("/register", handler.Authorize)
		api.POST("/login/", handler.Authenticate)
		// -----------------
		api.POST("/submitQuiz/", handler.StoreStudentQuiz)
		api.POST("/storeQnA/", handler.StoreQnA)
		api.GET("/loadQuiz", handler.LoadQuiz)
		// api.POST("/board/", handler.GetBoardData)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":1000")
}
