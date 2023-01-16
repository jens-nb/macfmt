package main

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		input    string
		testSize int
		want     []string
	}{
		{"abcdef123456", 1, []string{"a", "b", "c", "d", "e", "f", "1", "2", "3", "4", "5", "6"}},
		{"abcdef123456", 2, []string{"ab", "cd", "ef", "12", "34", "56"}},
		{"abcdef123456", 10, nil},
		{"a", 1, []string{"a"}},
		{"", 1, nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, _ := chunk(tt.input, tt.testSize)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
