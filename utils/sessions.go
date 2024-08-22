package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions/redis"
)

type SessionClaims struct {
	SessionID string
	Email     string
	jwt.StandardClaims
}

type SessionData struct {
	SessionID string
	Email     string
}

type claimsConstant struct {
	expiresAt     int64
	sessionSecret []byte
}

var Constant = &claimsConstant{
	expiresAt:     time.Now().Add(24 * time.Hour).Unix(),
	sessionSecret: []byte(os.Getenv("SESSION_SECRET")),
}

var redisKey string = os.Getenv("REDIS_KEY")
var redisPass string = os.Getenv("REDIS_PASS")

// function to initialize redis store.
func InitRedisStore(address string) (redis.Store, error) {
	store, err := redis.NewStore(10, "tcp", address, redisPass, []byte(redisKey))
	if err != nil {
		return nil, err
	}
	return store, nil
}

func GenerateJWTSession(sessionId, email string) (string, error) {
	//build new claims for JWT token.
	sessionClaims := &SessionClaims{
		SessionID: sessionId,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Constant.expiresAt,
		},
	}

	//build new token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims)

	//sign new token.
	signedToken, err := token.SignedString(Constant.sessionSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyJWTSession(token string) (*SessionData, bool, error) {
	tokenClaims := &SessionClaims{}

	parsed_token, err := jwt.ParseWithClaims(token, tokenClaims, func(tk *jwt.Token) (interface{}, error) {
		return Constant.sessionSecret, nil
	})
	if err != nil {
		return nil, false, errors.New("failed to parse token")
	}

	data := &SessionData{
		SessionID: tokenClaims.SessionID,
		Email:     tokenClaims.Email,
	}

	if parsed_token.Valid {
		return data, true, nil
	}
	return nil, false, errors.New("token is not valid")
}
