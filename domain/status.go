package domain

import "fmt"

var allowedTransitions = map[string]map[string]bool{"pending": {"registered": true, "cancelled": true}, "registered": {"approved": true, "cancelled": true}, "approved": {"processed": true, "cancelled": true}, "processed": {"notified": true}, "notified": {"attended": true, "declined": true}}

func ValidateTransition(from, to string) error {
	if from == to {
		return nil
	}
	next, ok := allowedTransitions[from]
	if !ok || !next[to] {
		return fmt.Errorf("transition %s to %s is not allowed", from, to)
	}
	return nil
}
func Statuses() []string {
	return []string{"pending", "registered", "approved", "processed", "notified", "attended", "declined", "cancelled"}
}
func IsKnownStatus(v string) bool {
	for _, s := range Statuses() {
		if s == v {
			return true
		}
	}
	return false
}
func TransitionPath(from, to string) []string {
	if from == to {
		return []string{from}
	}
	seen := map[string]bool{from: true}
	queue := [][]string{{from}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		tail := path[len(path)-1]
		for n := range allowedTransitions[tail] {
			if seen[n] {
				continue
			}
			next := append(append([]string{}, path...), n)
			if n == to {
				return next
			}
			seen[n] = true
			queue = append(queue, next)
		}
	}
	return nil
}
func ExplainStatus(v string) string {
	switch v {
	case "pending":
		return "awaiting registration"
	case "registered":
		return "awaiting review"
	case "approved":
		return "approved for processing"
	case "processed":
		return "ready for notification"
	case "notified":
		return "notification sent"
	case "attended":
		return "attendance confirmed"
	case "declined":
		return "declined by participant"
	case "cancelled":
		return "cancelled"
	}
	return "unknown"
}
