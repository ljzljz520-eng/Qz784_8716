package api

import (
	"context"
	"encoding/json"
	"frontend_go/domain"
	"frontend_go/workflow"
	"net/http"
)

type Handler struct {
	intake *workflow.Intake
	query  *workflow.QueryWorkflow
}

func NewHandler(i *workflow.Intake, q *workflow.QueryWorkflow) *Handler {
	return &Handler{intake: i, query: q}
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/health":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	case "/records":
		h.records(w, r)
	default:
		http.NotFound(w, r)
	}
}
func (h *Handler) records(w http.ResponseWriter, r *http.Request) {
	q := domain.Query{MeetingID: r.URL.Query().Get("meeting"), Status: r.URL.Query().Get("status")}
	page, e := h.query.Records(r.Context(), q)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(page)
}
func (h *Handler) Import(ctx context.Context, r domain.Record, p []domain.Participant) error {
	_, err := h.intake.Run(ctx, r, p)
	return err
}
