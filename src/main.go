package main

import (
	"net/http"

	"github.com/arx-8/try-go-gin/src/handlers"
	"github.com/arx-8/try-go-gin/src/middleware"
	"github.com/arx-8/try-go-gin/src/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.RecordUaAndTime)

	{
		booksRoute := r.Group("books")
		booksHandler := handlers.NewBooksHandler(service.BookService{})
		booksRoute.GET("", booksHandler.GetList)
		booksRoute.POST("", booksHandler.PostNew)
		booksRoute.GET("/:id", booksHandler.GetByID)
		booksRoute.PUT("/:id", booksHandler.Put)
		booksRoute.DELETE("/:id", booksHandler.Delete)
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.Done()
	})

	r.GET("/auth", auth)

	r.Run(":8080")
}

func auth(c *gin.Context) {
	accessToken := c.GetHeader("authorization")

	sess := session.Must(session.NewSession(aws.NewConfig().WithRegion("ap-northeast-1")))
	cognitoIDP := cognitoidentityprovider.New(sess)

	user, err := cognitoIDP.GetUser(&cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Username,
		"data": user.UserAttributes,
	})
}
