package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"frontend_go/domain"
	"go.etcd.io/bbolt"
	"sync"
)

var buckets = map[string][]byte{"records": []byte("records"), "users": []byte("users"), "events": []byte("events"), "audits": []byte("audits"), "meetings": []byte("meetings"), "participants": []byte("participants")}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func (s *Store) Path() string         { return s.path }
func encode(v any) ([]byte, error)    { return json.Marshal(v) }
func decode(data []byte, v any) error { return json.Unmarshal(data, v) }
func (s *Store) put(bucket, key string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	data, err := encode(v)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		raw := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if raw == nil {
			return fmt.Errorf("%s %s not found", bucket, key)
		}
		return decode(raw, v)
	})
}
func (s *Store) del(bucket, key string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Delete([]byte(key)) })
}
func (s *Store) PutRecord(v domain.Record) error { return s.put("records", v.ID, v) }
func (s *Store) GetRecord(id string) (domain.Record, error) {
	var v domain.Record
	err := s.get("records", id, &v)
	return v, err
}
func (s *Store) DeleteRecord(id string) error { return s.del("records", id) }
func (s *Store) PutUser(v domain.User) error  { return s.put("users", v.ID, v) }
func (s *Store) GetUser(id string) (domain.User, error) {
	var v domain.User
	err := s.get("users", id, &v)
	return v, err
}
func (s *Store) PutEvent(v domain.Event) error     { return s.put("events", v.ID, v) }
func (s *Store) PutAudit(v domain.Audit) error     { return s.put("audits", v.ID, v) }
func (s *Store) PutMeeting(v domain.Meeting) error { return s.put("meetings", v.ID, v) }
func (s *Store) GetMeeting(id string) (domain.Meeting, error) {
	var v domain.Meeting
	err := s.get("meetings", id, &v)
	return v, err
}
func (s *Store) PutParticipant(v domain.Participant) error { return s.put("participants", v.ID, v) }
func (s *Store) GetParticipant(id string) (domain.Participant, error) {
	var v domain.Participant
	err := s.get("participants", id, &v)
	return v, err
}
