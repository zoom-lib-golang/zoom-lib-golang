// +build integration

package zoom

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
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

func TestCreateDeleteUser(t *testing.T) {
	var (
		apiKey      = os.Getenv("ZOOM_API_KEY")
		apiSecret   = os.Getenv("ZOOM_API_SECRET")
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
	)

	APIKey = apiKey
	APISecret = apiSecret

	ms := time.Now().Unix() * int64(time.Microsecond)
	email := strings.Replace(primaryUser, "@", fmt.Sprintf("%d@", ms), 1)

	createOpts := CreateUserOptions{
		Action: CustCreate,
		UserInfo: CreateUserInfo{
			Type:  Basic,
			Email: email,
		},
	}
	user, err := CreateUser(createOpts)
	if err != nil {
		t.Fatalf("got error creating user: %+v\n", err)
	}

	if user.Email != email {
		t.Fatalf("expected %s, got %s\n", email, user.Email)
	}

	deleteOpts := DeleteUserOptions{
		EmailOrID: user.Email,
		Action:    DeleteAction,
	}
	err = DeleteUser(deleteOpts)
	if err != nil {
		t.Fatalf("got error deleting user: %+v\n", err)
	}
}
