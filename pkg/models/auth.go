package models

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"time"

	"go-web-server/lib"

	"golang.org/x/crypto/bcrypt"
)

type Authorization struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    AuthToken
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewAuthToken struct {
	token string
	data  interface{}
}

type AuthToken struct {
	Token string `json:"token"`
}

func (auth *Authorization) GetValidUser() (User, error) {
	db := lib.GetDB()
	credentials := Credentials{Email: auth.Email, Password: auth.Password}
	var user User
	err := db.Where("email = ?", credentials.Email).First(&user).Error
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(credentials.Password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (auth *Authorization) CreateSession(data interface{}) AuthToken {
	var authToken AuthToken
	authToken.Token = genSessionId(32)
	newToken := NewAuthToken{token: authToken.Token, data: data}
	err := newToken.SaveToken()
	if err != nil {
		log.Fatal(err)
	}
	return authToken
}

func (newToken *NewAuthToken) SaveToken() error {
	redis := lib.RedisCache{}
	key := newToken.createTokenKey()
	_, err := redis.Set(key, newToken.data, time.Hour).Result()
	return err
}

func (newToken *NewAuthToken) createTokenKey() string {
	return "app:token:" + newToken.token
}

func (auth *AuthToken) GenereteToken() string {
	token := genSessionId(32)
	return token
}

func genSessionId(n int) string {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
