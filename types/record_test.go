package types

import "testing"

func TestPostResults(t *testing.T) {
	testRecord := NewRecord(*NewWord("testWord"), *NewUser("testUser"), 0, "t,e,w")
	testRecord.PostResults("")
}
