package storage

import (
	"frontend_go/domain"
	"os"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	p := "store-test.db"
	defer os.Remove(p)
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := domain.NewRecord("r", "m", "p")
	if e = s.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	got, e := s.GetRecord("r")
	if e != nil || got.ID != "r" {
		t.Fatal(e)
	}
}
