package service

import (
	"context"
	"fmt"
	"frontend_go/domain"
	"frontend_go/storage"
	"time"
)

type ReviewService struct {
	store    *storage.Store
	notifier func(domain.Record) error
}

func NewReviewService(s *storage.Store, n func(domain.Record) error) *ReviewService {
	if n == nil {
		n = func(domain.Record) error { return nil }
	}
	return &ReviewService{store: s, notifier: n}
}
func (s *ReviewService) Approve(ctx context.Context, id, actor string) (domain.Record, error) {
	return s.move(ctx, id, "approved", actor)
}
func (s *ReviewService) Process(ctx context.Context, id, actor string) (domain.Record, error) {
	return s.move(ctx, id, "processed", actor)
}
func (s *ReviewService) move(ctx context.Context, id, to, actor string) (domain.Record, error) {
	r, err := s.store.GetRecord(id)
	if err != nil {
		return r, err
	}
	if err = ctx.Err(); err != nil {
		return r, err
	}
	if err = domain.ValidateTransition(r.Status, to); err != nil {
		return r, err
	}
	before := r.Status
	r.Status = to
	r.Version++
	r.UpdatedAt = time.Now()
	if err = s.store.PutRecord(r); err != nil {
		return r, err
	}
	a := domain.Audit{ID: fmt.Sprintf("%s-%d", id, r.Version), Entity: "Record", EntityID: id, Action: to, Before: before, After: to, Actor: actor, At: time.Now()}
	if err = s.store.PutAudit(a); err != nil {
		return r, err
	}
	return r, nil
}
func (s *ReviewService) Notify(ctx context.Context, id string) (domain.Record, error) {
	r, err := s.move(ctx, id, "notified", "system")
	if err != nil {
		return r, err
	}
	if err = s.notifier(r); err != nil {
		return r, err
	}
	return r, nil
}
