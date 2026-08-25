package reporting

import (
	"encoding/json"
	"fmt"
	"frontend_go/domain"
)

func ExportJSON(records []domain.Record) (string, error) {
	b, e := json.MarshalIndent(records, "", "  ")
	return string(b), e
}
func ExportCSV(records []domain.Record) string {
	out := "id,meeting,participant,status,version\n"
	for _, r := range records {
		out += fmt.Sprintf("%s,%s,%s,%s,%d\n", r.ID, r.MeetingID, r.ParticipantID, r.Status, r.Version)
	}
	return out
}
func CountByMeeting(records []domain.Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[r.MeetingID]++
	}
	return out
}
