// +build integration

package zoom

import (
	"os"
	"testing"
)

func TestListUsers(t *testing.T) {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
	)

	APIKey = apiKey
	APISecret = apiSecret

	_, err := ListUsers(ListUsersOptions{
		PageSize:   1,
		PageNumber: 1,
	})
	if err != nil {
		t.Fatalf("got error listing users: %+v\n", err)
	}
}

func TestGetUser(t *testing.T) {
	var (
		apiKey      = os.Getenv("ZOOM_API_KEY")
		apiSecret   = os.Getenv("ZOOM_API_SECRET")
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
	)

	APIKey = apiKey
	APISecret = apiSecret

	user, err := GetUser(GetUserOpts{EmailOrID: primaryUser})
	if err != nil {
		t.Fatalf("got error listing users: %+v\n", err)
	}

	if user.Email != primaryUser {
		t.Fatalf("expected %s, got %s\n", primaryUser, user.Email)
	}
}
