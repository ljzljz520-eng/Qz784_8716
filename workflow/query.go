package workflow

import (
	"context"
	"frontend_go/domain"
	"frontend_go/service"
)

type QueryWorkflow struct {
	reg  *service.RegistrationService
	tree *service.TreeService
}

func NewQueryWorkflow(r *service.RegistrationService, t *service.TreeService) *QueryWorkflow {
	return &QueryWorkflow{reg: r, tree: t}
}
func (w *QueryWorkflow) Records(ctx context.Context, q domain.Query) (domain.Page, error) {
	return w.reg.Find(ctx, q)
}
func (w *QueryWorkflow) Tree(ctx context.Context, meeting string) ([]domain.Participant, error) {
	root, e := w.tree.Build(ctx, meeting)
	if e != nil {
		return nil, e
	}
	return w.tree.Flatten(root), nil
}
func (w *QueryWorkflow) Status(ctx context.Context, id string) (string, error) {
	r, e := w.reg.Find(ctx, domain.Query{})
	if e != nil {
		return "", e
	}
	for _, v := range r.Records {
		if v.ID == id {
			return v.Status, nil
		}
	}
	return "", context.Canceled
}
