package workflow

import (
	"context"
	"frontend_go/domain"
	"frontend_go/service"
)

type Intake struct {
	reg  *service.RegistrationService
	tree *service.TreeService
}

func NewIntake(r *service.RegistrationService, t *service.TreeService) *Intake {
	return &Intake{reg: r, tree: t}
}
func (w *Intake) Run(ctx context.Context, r domain.Record, people []domain.Participant) (domain.Record, error) {
	if err := ctx.Err(); err != nil {
		return r, err
	}
	if err := w.tree.Import(ctx, r.MeetingID, people); err != nil {
		return r, err
	}
	return w.reg.Register(ctx, r)
}
func (w *Intake) Validate(r domain.Record) error { return domain.ValidateRecord(r) }
func (w *Intake) Current(ctx context.Context, q domain.Query) (domain.Page, error) {
	return w.reg.Find(ctx, q)
}
