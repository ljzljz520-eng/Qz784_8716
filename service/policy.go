package service

import "frontend_go/domain"

func Decision1(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible1(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision2(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible2(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision3(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible3(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision4(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible4(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision5(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible5(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision6(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible6(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision7(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible7(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision8(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible8(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision9(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible9(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision10(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible10(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision11(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible11(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision12(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible12(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision13(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible13(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision14(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible14(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision15(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible15(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision16(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible16(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision17(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible17(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision18(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible18(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision19(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible19(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision20(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible20(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision21(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible21(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision22(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible22(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision23(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible23(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision24(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible24(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision25(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible25(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision26(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible26(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision27(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible27(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision28(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible28(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision29(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible29(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision30(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible30(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision31(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible31(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision32(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible32(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision33(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible33(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision34(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible34(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
func Decision35(r domain.Record) string {
	if r.IsTerminal() {
		return "terminal"
	}
	if r.Status == "pending" {
		return "waiting"
	}
	if r.Status == "registered" {
		return "review"
	}
	return "progress"
}
func Eligible35(r domain.Record) bool {
	if !domain.IsKnownStatus(r.Status) {
		return false
	}
	return !r.IsTerminal()
}
