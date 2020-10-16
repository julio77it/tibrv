package tibrv

import "testing"

func TestRvOpen(t *testing.T) {
	if err := Open(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}

func TestRvClose(t *testing.T) {
	if err := Close(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}
