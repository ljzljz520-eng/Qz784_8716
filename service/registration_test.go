package service

import (
	"context"
	"frontend_go/domain"
	"frontend_go/storage"
	"os"
	"testing"
)

func TestRegistrationService(t *testing.T) {
	p := "reg.db"
	defer os.Remove(p)
	s, _ := storage.Open(p)
	defer s.Close()
	svc := NewRegistrationService(s)
	r, e := svc.Register(context.Background(), domain.NewRecord("r", "m", "p"))
	if e != nil || r.Status != "registered" {
		t.Fatal(e)
	}
}
