package service

import (
	"context"
	"fmt"
	"frontend_go/domain"
	"frontend_go/storage"
	"time"
)

type RegistrationService struct {
	store *storage.Store
	clock func() time.Time
}

func NewRegistrationService(s *storage.Store) *RegistrationService {
	return &RegistrationService{store: s, clock: time.Now}
}
func (s *RegistrationService) Register(ctx context.Context, r domain.Record) (domain.Record, error) {
	if err := domain.ValidateRecord(r); err != nil {
		return r, err
	}
	if err := ctx.Err(); err != nil {
		return r, err
	}
	old, err := s.store.GetRecord(r.ID)
	if err == nil {
		r.Version = old.Version + 1
		if e := domain.ValidateTransition(old.Status, "registered"); e != nil {
			return r, e
		}
	}
	r.Status = "registered"
	r.UpdatedAt = s.clock()
	if err := s.store.PutRecord(r); err != nil {
		return r, err
	}
	return r, nil
}
func (s *RegistrationService) Cancel(ctx context.Context, id string) error {
	r, err := s.store.GetRecord(id)
	if err != nil {
		return err
	}
	if err = domain.ValidateTransition(r.Status, "cancelled"); err != nil {
		return err
	}
	if err = ctx.Err(); err != nil {
		return err
	}
	r.Status = "cancelled"
	r.Version++
	r.UpdatedAt = s.clock()
	return s.store.PutRecord(r)
}
func (s *RegistrationService) Find(ctx context.Context, q domain.Query) (domain.Page, error) {
	if err := ctx.Err(); err != nil {
		return domain.Page{}, err
	}
	return s.store.ListRecords(q)
}
func (s *RegistrationService) Describe(id string) (string, error) {
	r, e := s.store.GetRecord(id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s:%s:v%d", r.ID, domain.ExplainStatus(r.Status), r.Version), nil
}
