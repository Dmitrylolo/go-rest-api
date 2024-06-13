//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func createToken() string {
	tokewn := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := tokewn.SignedString([]byte("my-secret-key"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("can create a new comment", func(t *testing.T) {
		client := resty.New()
		token := createToken()
		resp, err := client.R().
			SetBody(`{"slug":"slug","author":"author","body":"body"}`).
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", "Bearer "+token).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err, "Error while creating a new comment")

		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("can't create a new comment without JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug":"slug","author":"author","body":"body"}`).
			SetHeader("Content-Type", "application/json").
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err, "Error while creating a new comment")

		assert.Equal(t, 401, resp.StatusCode())
	})
}
