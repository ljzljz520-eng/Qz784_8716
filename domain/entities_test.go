package domain

import (
	"testing"
	"time"
)

func TestEntityValidation(t *testing.T) {
	r := NewRecord("r", "m", "p")
	if ValidateRecord(r) != nil {
		t.Fatal("record")
	}
	if ValidateMeeting(NewMeeting("m", "Meeting 3", "u", time.Now())) != nil {
		t.Fatal("meeting")
	}
}
