package domain

import "testing"

func TestStatusRules(t *testing.T) {
	if ValidateTransition("pending", "approved") == nil {
		t.Fatal("invalid transition")
	}
	if len(TransitionPath("pending", "processed")) != 4 {
		t.Fatal("path")
	}
}
