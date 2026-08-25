package storage

import (
	"frontend_go/domain"
	"os"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "reopen.db"
	defer os.Remove(p)
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.PutRecord(domain.NewRecord("persist", "m", "p")); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
}
