package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordUaAndTime(c *gin.Context) {
	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// oldTime := time.Now()
	// ua := c.GetHeader("user-agent")

	// c.Next()

	// logger.Info(
	// 	"Got request",
	// 	zap.String("path", c.Request.URL.Path),
	// 	zap.String("ua", ua),
	// 	zap.Int("status", c.Writer.Status()),
	// 	zap.Duration("elapsed", time.Since(oldTime)),
	// )
}

func AuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	// validate value
	r := regexp.MustCompile(`^Bearer e`)
	if strings.Count(authorization, ".") != 2 || !r.MatchString(authorization) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization header value",
		})
		return
	}

	// parse
	jwtAccessToken := strings.Replace(authorization, "Bearer ", "", 1)
	parsedJWT := strings.Split(jwtAccessToken, ".")
	decodedPayload, _ := base64url.Decode(parsedJWT[1])
	payload := AccessTokenPayload{}
	json.Unmarshal(decodedPayload, &payload)

	if err := payload.Check(); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()

	logger, _ := zap.NewProduction()
	ua := c.GetHeader("user-agent")
	logger.Info(
		"Got request",
		zap.String("path", c.Request.RequestURI),
		zap.String("ua", ua),
		zap.Int("status", c.Writer.Status()),
		zap.String("requester_user_id", payload.Sub),
	)
}

type AccessTokenPayload struct {
	AuthTime uint `json:"auth_time"`

	// client_id for cognito-idp
	ClientID string `json:"client_id"`
	EventID  string `json:"event_id"`

	// Expiration date of the token
	Exp uint `json:"exp"`
	Iat uint `json:"iat"`

	// Identifier of the token issuer
	Iss       string `json:"iss"`
	Jti       string `json:"jti"`
	OriginJti string `json:"origin_jti"`
	Scope     string `json:"scope"`

	// User identifier
	Sub      string `json:"sub"`
	TokenUse string `json:"token_use"`
	Username string `json:"username"`
}

func (a AccessTokenPayload) isValidIss() bool {
	REGION := "ap-northeast-1"
	USER_POOL_ID := "ap-northeast-1_4CFdnlqYG"
	return a.Iss == "https://cognito-idp."+REGION+".amazonaws.com/"+USER_POOL_ID
}

func (a AccessTokenPayload) isValidClientID() bool {
	// REGION := "ap-northeast-1"
	// USER_POOL_ID := "ap-northeast-1_4CFdnlqYG"
	USER_POOL_WEB_CLIENT_ID := "2p8i215357kcml1if0blmto071"
	return a.ClientID == USER_POOL_WEB_CLIENT_ID
}

func (a AccessTokenPayload) isExpired() bool {
	return int64(a.Exp) < time.Now().Unix()
}

func (a AccessTokenPayload) Check() error {
	var msgs []string

	if !a.isValidIss() {
		msgs = append(msgs, "Invalid iss")
	}
	if !a.isValidClientID() {
		msgs = append(msgs, "Invalid client_id")
	}
	if a.isExpired() {
		msgs = append(msgs, "Expired")
	}

	if len(msgs) == 0 {
		return nil
	}

	return errors.New(strings.Join(msgs, ", "))
}
