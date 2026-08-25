package storage

import (
	"errors"
	"frontend_go/domain"
	"go.etcd.io/bbolt"
	"sort"
)

var ErrClosed = errors.New("store closed")

func (s *Store) ListRecords(q domain.Query) (domain.Page, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return domain.Page{}, ErrClosed
	}
	all := []domain.Record{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r domain.Record
			if e := decode(v, &r); e != nil {
				return e
			}
			if q.MeetingID != "" && r.MeetingID != q.MeetingID {
				return nil
			}
			if q.Status != "" && r.Status != q.Status {
				return nil
			}
			all = append(all, r)
			return nil
		})
	})
	if err != nil {
		return domain.Page{}, err
	}
	sort.Slice(all, func(i, j int) bool { return all[i].UpdatedAt.Before(all[j].UpdatedAt) })
	limit, offset := q.NormalizedLimit(), q.NormalizedOffset()
	total := len(all)
	if offset > total {
		offset = total
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return domain.Page{Records: all[offset:end], Total: total, Limit: limit, Offset: offset}, nil
}
func (s *Store) ListParticipants(meeting string) ([]domain.Participant, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, ErrClosed
	}
	out := []domain.Participant{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("participants")).ForEach(func(_, v []byte) error {
			var p domain.Participant
			if e := decode(v, &p); e != nil {
				return e
			}
			if p.MeetingID == meeting {
				out = append(out, p)
			}
			return nil
		})
	})
	return out, err
}
func (s *Store) ListAudits(entityID string) ([]domain.Audit, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, ErrClosed
	}
	out := []domain.Audit{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("audits")).ForEach(func(_, v []byte) error {
			var a domain.Audit
			if e := decode(v, &a); e != nil {
				return e
			}
			if a.EntityID == entityID {
				out = append(out, a)
			}
			return nil
		})
	})
	return out, err
}
