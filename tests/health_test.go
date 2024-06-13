//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func TestHealthCheckEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/alive")
	assert.NoError(t, err, "Error while calling health check endpoint")

	assert.Equal(t, 200, resp.StatusCode())
}
