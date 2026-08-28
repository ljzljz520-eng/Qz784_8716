package workflow

import "frontend_go/domain"

type Checklist struct {
	Name     string
	Required []string
}

func (w *Checklist) Step1(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step2(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step3(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step4(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step5(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step6(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step7(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step8(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step9(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step10(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step11(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step12(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step13(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step14(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step15(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step16(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step17(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step18(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step19(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step20(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step21(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step22(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step23(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step24(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step25(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step26(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step27(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step28(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step29(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step30(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step31(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step32(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step33(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step34(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
func (w *Checklist) Step35(r domain.Record) bool {
	if r.ID == "" {
		return false
	}
	if r.Status == "cancelled" {
		return false
	}
	return true
}
