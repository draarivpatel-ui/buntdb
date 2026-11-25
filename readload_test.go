package buntdb

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestReadLoadMalformedCommands(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"short set command", "*3\r\n$1\r\nS\r\n$1\r\nk\r\n$1\r\nv\r\n"},
		{"short del command", "*2\r\n$2\r\nDE\r\n$1\r\nk\r\n"},
		{"empty command", "*1\r\n$0\r\n\r\n"},
		{"part size overflow", "*3\r\n$3\r\nset\r\n$9223372036854775808\r\n"},
		{"command count overflow", "*9223372036854775808\r\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{}
			_, err := db.readLoad(strings.NewReader(tt.input), time.Now())
			if !errors.Is(err, ErrInvalid) {
				t.Fatalf("expected ErrInvalid, got %v", err)
			}
		})
	}
}
