package functional

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/pkg/hangle"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var (
	base_url = "http://127.0.0.1:8080"
)

func reset() error {
	db, err := database.Open("localhost")
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		return retErr
	}

	db.Migrator().DropTable(&hangle.Record{})
	db.Migrator().DropTable(&hangle.User{})
	db.Migrator().DropTable(&hangle.Word{})

	logrus.Debug("Automigrating database")
	err = db.AutoMigrate(&hangle.User{}, &hangle.Word{}, &hangle.Record{})
	if err != nil {
		return fmt.Errorf("unable to automigrate: %w", err)
	}

	return nil
}

// TestCreateWord seeds a word and checks if it was created
func TestCreateUser(t *testing.T) {
	client := hangle.NewClient(hangle.NewConfig(base_url), http.DefaultClient)
	{
		require.NoError(t, reset())
		u := hangle.NewUser("foo")
		_, err := client.PostUser(u)
		require.NoError(t, err)
		u = hangle.NewUser("bar")
		_, err = client.PostUser(u)
		require.NoError(t, err)
	}
	{
		require.NoError(t, reset())
		u := hangle.NewUser("foo")
		_, err := client.PostUser(u)
		require.NoError(t, err)
		expectedErr := "ERROR: duplicate key value violates unique constraint \\\"users_username_key\\\""
		_, err = client.PostUser(u)
		require.ErrorContains(t, err, expectedErr)
	}
}

func TestGetAllUsers(t *testing.T) {
	client := hangle.NewClient(hangle.NewConfig(base_url), http.DefaultClient)
	{
		users := []string{"user1", "user2", "user3"}
		err := reset()
		require.NoError(t, err)
		for _, user := range users {
			u := hangle.NewUser(user)
			client.PostUser(u)
			require.NoError(t, err)
		}
		userGet, err := client.GetAllUsers()
		require.NoError(t, err)
		require.Equal(t, len(users), len(userGet))
	}
}
