package tibrv

import (
	"testing"
)

func TestNewRvError(t *testing.T) {
	err := NewRvError(5)
	if err == nil {
		t.Fatalf("Expected 5, got nil")
	}
	if err.Code != 5 {
		t.Fatalf("Expected 5, got %v", err)
	}
	if err.Text != "Arguments conflict" {
		t.Fatalf("Expected 'Arguments conflict', got %s", err.Text)
	}
	if err.String() != "5 - Arguments conflict" {
		t.Fatalf("Expected '5 - Arguments conflict', got %s", err.Text)
	}
	err = NewRvError(0)
	if err != nil {
		t.Fatalf("Expected nil, got nil")
	}
}
