package main

import (
	"context"
	"frontend_go/api"
	"frontend_go/service"
	"frontend_go/storage"
	"frontend_go/workflow"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	path := "meeting.db"
	if v := os.Getenv("MEETING_DB"); v != "" {
		path = v
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	reg := service.NewRegistrationService(s)
	tree := service.NewTreeService(s)
	review := service.NewReviewService(s, nil)
	h := api.NewHandler(workflow.NewIntake(reg, tree), workflow.NewQueryWorkflow(reg, tree))
	_ = review
	server := &http.Server{Addr: ":8080", Handler: h}
	go func() {
		if e := server.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			log.Print(e)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	server.Shutdown(context.Background())
}
