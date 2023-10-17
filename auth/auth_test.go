package auth_test

import (
	"testing"

	"github.com/kisukegremory/plateapi/auth"
)

func TestGenerateToken(t *testing.T) {
	_, err := auth.GenerateJwt()
	if err != nil {
		t.Error("Problems on token generation")
	}
}
