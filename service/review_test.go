package service

import (
	"context"
	"frontend_go/domain"
	"frontend_go/storage"
	"os"
	"testing"
)

func TestReviewService(t *testing.T) {
	p := "review.db"
	defer os.Remove(p)
	s, _ := storage.Open(p)
	defer s.Close()
	r := domain.NewRecord("r", "m", "p")
	r.Status = "registered"
	s.PutRecord(r)
	svc := NewReviewService(s, nil)
	if _, e := svc.Approve(context.Background(), "r", "u"); e != nil {
		t.Fatal(e)
	}
}
