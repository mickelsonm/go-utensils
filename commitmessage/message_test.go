package commitmessage

import (
	"testing"
)

func TestCommitMessage(t *testing.T) {
	t.Log("Testing internal fetch")
	mess, err := fetchCommitMessage()
	if err != nil {
		t.Log("Encounterd an error fetching a message")
		t.Fatal(err)
	}
	if mess == "" {
		t.Log("message shouldn't be empty")
		t.Fatal(err)
	}

	t.Log("Testing creating a new message")
	msg, err := NewMessage()
	if err != nil {
		t.Log("Encountered an error while getting a message")
		t.Fatal(err)
	}

	t.Log("Testing text output")
	_, err = msg.ToString()
	if err != nil {
		t.Log("Encountered an error while displaying text output")
		t.Fatal(err)
	}

	t.Log("Testing json output")
	_, err = msg.ToJSON()
	if err != nil {
		t.Log("Encountered an error while displaying json output")
		t.Fatal(err)
	}

	t.Log("Testing xml output")
	_, err = msg.ToXML()
	if err != nil {
		t.Log("Encountered an error while displaying xml output")
		t.Fatal(err)
	}
}
