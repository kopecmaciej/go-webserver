package models_test

import (
	"go-web-server/pkg/models"
	"reflect"
	"testing"
)

func TestCreateToken(t *testing.T) {
	var auth models.Authorization
	token := auth.Token.GenereteToken()
	tokenType := reflect.TypeOf(token).Kind()
	if tokenType.String() != "string" {
		t.Fail()
	}
	tokenLength := len(token)
	if tokenLength != 44 {
		t.Fail()
	}
}
