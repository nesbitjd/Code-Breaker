package types

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type recordClient struct {
	Error error
}

func (r *recordClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, r.Error
}

func TestPostResults(t *testing.T) {
	{
		httpClient := recordClient{Error: nil}
		c := NewClient(NewConfig("localhost:8878"), &httpClient)
		testRecord := NewRecord(*NewWord("testWord"), *NewUser("testUser"), 0, "t,e,w")
		_, err := c.PostRecord(testRecord)
		require.NoError(t, err)
	}
	{
		httpClient := recordClient{Error: fmt.Errorf("test error")}
		c := NewClient(NewConfig("localhost:8878"), &httpClient)
		testRecord := NewRecord(*NewWord("testWord"), *NewUser("testUser"), 0, "t,e,w")
		_, err := c.PostRecord(testRecord)
		require.ErrorContains(t, err, "unable to do http request")

	}
}

func TestDeleteResults(t *testing.T) {

}
