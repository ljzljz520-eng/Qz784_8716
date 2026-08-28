package domain

import "time"

type Record struct {
	ID            string
	MeetingID     string
	ParticipantID string
	Status        string
	Version       int
	UpdatedAt     time.Time
	Notes         string
}
type User struct {
	ID        string
	Name      string
	Email     string
	Role      string
	Active    bool
	CreatedAt time.Time
}
type Event struct {
	ID       string
	RecordID string
	Kind     string
	Payload  string
	At       time.Time
	Actor    string
}
type Audit struct {
	ID       string
	Entity   string
	EntityID string
	Action   string
	Before   string
	After    string
	At       time.Time
	Actor    string
}
type Meeting struct {
	ID        string
	Title     string
	Date      time.Time
	Owner     string
	State     string
	CreatedAt time.Time
}
type Participant struct {
	ID        string
	MeetingID string
	ParentID  string
	Name      string
	Email     string
	State     string
	Depth     int
}
type TreeNode struct {
	Participant Participant
	Children    []*TreeNode
}
type Query struct {
	MeetingID string
	Status    string
	Text      string
	Limit     int
	Offset    int
}
type Page struct {
	Records []Record
	Total   int
	Limit   int
	Offset  int
}

func NewRecord(id, meeting, participant string) Record {
	return Record{ID: id, MeetingID: meeting, ParticipantID: participant, Status: "pending", Version: 1, UpdatedAt: time.Now()}
}
func NewMeeting(id, title, owner string, date time.Time) Meeting {
	return Meeting{ID: id, Title: title, Owner: owner, Date: date, State: "draft", CreatedAt: time.Now()}
}
func NewParticipant(id, meeting, parent, name, email string) Participant {
	return Participant{ID: id, MeetingID: meeting, ParentID: parent, Name: name, Email: email, State: "invited"}
}
func (r Record) IsCurrent(v int) bool { return r.Version == v }
func (r Record) IsTerminal() bool {
	return r.Status == "cancelled" || r.Status == "attended" || r.Status == "declined"
}
func (m Meeting) IsOpen() bool          { return m.State == "draft" || m.State == "active" }
func (p Participant) IsRoot() bool      { return p.ParentID == "" }
func (p Participant) CanRegister() bool { return p.State == "invited" || p.State == "registered" }
func (q Query) NormalizedLimit() int {
	if q.Limit <= 0 || q.Limit > 100 {
		return 20
	}
	return q.Limit
}
func (q Query) NormalizedOffset() int {
	if q.Offset < 0 {
		return 0
	}
	return q.Offset
}
