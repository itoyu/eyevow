package eyevow

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type accessTokenClaims struct {
	User primitive.ObjectID `json:"user"`
	jwt.StandardClaims
}

type auth struct {
	secret string
}

var defaultAuth = newAuth(defaultEnv.Secret)

func newAuth(sec string) *auth {
	return &auth{sec}
}

func (a *auth) Validate(tk string) primitive.ObjectID {
	token, err := jwt.ParseWithClaims(tk, &accessTokenClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})
	if err != nil {
		return primitive.ObjectID{}
	}

	if claims, ok := token.Claims.(*accessTokenClaims); ok && token.Valid {
		return claims.User
	}
	return primitive.ObjectID{}
}

func (a *auth) Publish(uid primitive.ObjectID) string {
	d := time.Duration(43800 * time.Hour)

	tk, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessTokenClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
		},
	}).SignedString([]byte(a.secret))

	if err != nil {
		panic(err)
	}

	return tk
}
