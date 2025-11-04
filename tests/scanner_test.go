package main

import (
	"testing"

	"PortScanner/internal/scanner"
)

func TestParsPortRange(t *testing.T) {
	start, end, err := scanner.ParsePortRange("20-25")
	if err != nil {
		t.Error("Unexpected error: %v", err)
	}
	if start != 20 || end != 25 {
		t.Error("Expected error 20-24, for %d-%d", start, end)
	}

	_, _, err = scanner.ParsePortRange("bad-input")
	if err == nil {
		t.Error("Expected error for bad input, for nil")
	}
}
