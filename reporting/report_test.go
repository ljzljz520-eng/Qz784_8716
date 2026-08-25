package reporting

import (
	"frontend_go/domain"
	"testing"
)

func TestReport(t *testing.T) {
	s := Summarize([]domain.Record{{Status: "pending"}, {Status: "cancelled"}})
	if s.Total != 2 || s.ByStatus["cancelled"] != 1 {
		t.Fatal(s)
	}
}
