package functional

import (
	"fmt"
	"testing"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"
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

	db.Migrator().DropTable(&types.Record{})
	db.Migrator().DropTable(&types.User{})
	db.Migrator().DropTable(&types.Word{})

	logrus.Debug("Automigrating database")
	err = db.AutoMigrate(&types.User{}, &types.Word{}, &types.Record{})
	if err != nil {
		return fmt.Errorf("unable to automigrate: %w", err)
	}

	return nil
}

// TestCreateWord seeds a word and checks if it was created
func TestCreateUser(t *testing.T) {
	{
		require.NoError(t, reset())
		require.NoError(t, types.NewUser("foo").PostUser(base_url))
		require.NoError(t, types.NewUser("bar").PostUser(base_url))
	}
	{
		require.NoError(t, reset())
		require.NoError(t, types.NewUser("foo").PostUser(base_url))
		expectedErr := "ERROR: duplicate key value violates unique constraint \"users_username_key\""
		require.ErrorContains(t, types.NewUser("foo").PostUser(base_url), expectedErr)
	}
}

func TestGetAllUsers(t *testing.T) {
	{
		users := []string{"user1", "user2", "user3"}
		err := reset()
		require.NoError(t, err)
		for _, u := range users {
			require.NoError(t, types.NewUser(u).PostUser(base_url))
		}
		userGet, err := types.GetAllUsers(base_url)
		require.NoError(t, err)
		require.Equal(t, len(users), len(userGet))
	}
}
