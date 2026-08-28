package workflow

import (
	"context"
	"frontend_go/domain"
	"frontend_go/service"
)

type Processing struct{ review *service.ReviewService }

func NewProcessing(r *service.ReviewService) *Processing { return &Processing{review: r} }
func (w *Processing) Run(ctx context.Context, id, actor string) (domain.Record, error) {
	r, e := w.review.Approve(ctx, id, actor)
	if e != nil {
		return r, e
	}
	return w.review.Process(ctx, r.ID, actor)
}
func (w *Processing) Cancel(ctx context.Context, id string) error { return nil }
func (w *Processing) Notify(ctx context.Context, id string) (domain.Record, error) {
	return w.review.Notify(ctx, id)
}
