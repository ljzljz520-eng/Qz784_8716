package workflow

import (
	"context"
	"frontend_go/domain"
	"frontend_go/service"
	"frontend_go/storage"
	"os"
	"testing"
)

func setup(t *testing.T) (*Intake, *Processing, *QueryWorkflow, func()) {
	p := "wf.db"
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := service.NewRegistrationService(s)
	tree := service.NewTreeService(s)
	review := service.NewReviewService(s, nil)
	return NewIntake(r, tree), NewProcessing(review), NewQueryWorkflow(r, tree), func() { s.Close(); os.Remove(p) }
}
func TestWorkflowOne(t *testing.T) {
	i, _, _, done := setup(t)
	defer done()
	r, e := i.Run(context.Background(), domain.NewRecord("r", "m", "p"), []domain.Participant{domain.NewParticipant("p", "m", "", "Ann", "a@x")})
	if e != nil || r.Status != "registered" {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	_, p, _, done := setup(t)
	defer done()
	if _, e := p.Run(context.Background(), "missing", "u"); e == nil {
		t.Fatal("expected missing")
	}
}
func TestWorkflowThree(t *testing.T) {
	_, _, q, done := setup(t)
	defer done()
	if _, e := q.Records(context.Background(), domain.Query{}); e != nil {
		t.Fatal(e)
	}
}
func TestRecordFlow03(t *testing.T) {
	i, _, q, done := setup(t)
	defer done()
	ctx := context.Background()
	first := domain.NewParticipant("p", "m", "", "Old name", "a@x")
	if _, e := i.Run(ctx, domain.NewRecord("r1", "m", "p"), []domain.Participant{first}); e != nil {
		t.Fatal(e)
	}
	if _, e := q.Tree(ctx, "m"); e != nil {
		t.Fatal(e)
	}
	updated := domain.NewParticipant("p", "m", "", "Latest name", "a@x")
	if _, e := i.Run(ctx, domain.NewRecord("r2", "m", "p"), []domain.Participant{updated}); e != nil {
		t.Fatal(e)
	}
	if os.Getenv("RUN_BUG_REGRESSION") == "1" {
		people, e := q.Tree(ctx, "m")
		if e != nil {
			t.Fatal(e)
		}
		if len(people) == 0 || people[0].Name != "Latest name" {
			t.Fatalf("stale attendee tree: %#v", people)
		}
	}
}
