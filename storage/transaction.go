package storage

import (
	"frontend_go/domain"
	"time"
)

type Transaction struct {
	store   *Store
	records []domain.Record
	events  []domain.Event
	audits  []domain.Audit
}

func (s *Store) Begin() *Transaction             { return &Transaction{store: s} }
func (t *Transaction) AddRecord(r domain.Record) { t.records = append(t.records, r) }
func (t *Transaction) AddEvent(e domain.Event)   { t.events = append(t.events, e) }
func (t *Transaction) AddAudit(a domain.Audit)   { t.audits = append(t.audits, a) }
func (t *Transaction) Commit() error {
	for _, r := range t.records {
		if err := t.store.PutRecord(r); err != nil {
			return err
		}
	}
	for _, e := range t.events {
		if err := t.store.PutEvent(e); err != nil {
			return err
		}
	}
	for _, a := range t.audits {
		if a.At.IsZero() {
			a.At = time.Now()
		}
		if err := t.store.PutAudit(a); err != nil {
			return err
		}
	}
	return nil
}
func (t *Transaction) Rollback() { t.records = nil; t.events = nil; t.audits = nil }
func (s *Store) Snapshot(ids []string) ([]domain.Record, error) {
	out := make([]domain.Record, 0, len(ids))
	for _, id := range ids {
		r, e := s.GetRecord(id)
		if e != nil {
			return nil, e
		}
		out = append(out, r)
	}
	return out, nil
}
