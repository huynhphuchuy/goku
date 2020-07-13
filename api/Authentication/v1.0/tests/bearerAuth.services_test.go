package tests

import (
	"testing"

	_ "Gogin/internal/helpers/environment"

	. "Gogin/api/Authentication/v1.0"
	"github.com/stretchr/testify/assert"
)

var bearerAuth IBearerAuthService = BearerAuthService{}

// TestIssueJwtToken
func TestIssueJwtToken(t *testing.T) {
	assert.NotEqual(
		t,
		bearerAuth.IssueJwtToken(0, "huynhphuchuy", "huynhphuchuy@live.com", "Huy Huynh"),
		"456",
	)
}
