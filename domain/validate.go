package domain

import (
	"errors"
	"strings"
)

func ValidateRecord(r Record) error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("record id required")
	}
	if strings.TrimSpace(r.MeetingID) == "" {
		return errors.New("meeting id required")
	}
	if strings.TrimSpace(r.ParticipantID) == "" {
		return errors.New("participant id required")
	}
	if !IsKnownStatus(r.Status) {
		return errors.New("unknown status")
	}
	if r.Version < 1 {
		return errors.New("version must be positive")
	}
	return nil
}
func ValidateUser(u User) error {
	if strings.TrimSpace(u.ID) == "" || strings.TrimSpace(u.Name) == "" {
		return errors.New("user identity required")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("email required")
	}
	return nil
}
func ValidateMeeting(m Meeting) error {
	if m.ID == "" || m.Title == "" {
		return errors.New("meeting identity required")
	}
	if m.Date.IsZero() {
		return errors.New("meeting date required")
	}
	return nil
}
func ValidateParticipant(p Participant) error {
	if p.ID == "" || p.MeetingID == "" || p.Name == "" {
		return errors.New("participant identity required")
	}
	if p.Depth < 0 {
		return errors.New("depth cannot be negative")
	}
	return nil
}
func CleanText(v string) string { return strings.Join(strings.Fields(v), " ") }
