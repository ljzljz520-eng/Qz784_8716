package api

import (
	"frontend_go/service"
	"frontend_go/storage"
	"frontend_go/workflow"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandlerHealth(t *testing.T) {
	p := "api.db"
	defer os.Remove(p)
	s, _ := storage.Open(p)
	defer s.Close()
	h := NewHandler(workflow.NewIntake(service.NewRegistrationService(s), service.NewTreeService(s)), workflow.NewQueryWorkflow(service.NewRegistrationService(s), service.NewTreeService(s)))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
