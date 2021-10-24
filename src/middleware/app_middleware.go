package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordUaAndTime(c *gin.Context) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	oldTime := time.Now()
	ua := c.GetHeader("user-agent")

	c.Next()

	logger.Info(
		"Got request",
		zap.String("path", c.Request.URL.Path),
		zap.String("ua", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Since(oldTime)),
	)
}
