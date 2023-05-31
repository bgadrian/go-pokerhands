package five

import (
	"testing"
)

func TestParseCard(t *testing.T) {
	tests := []struct {
		args    string
		wantErr bool
	}{
		{"Ad", false},
		{"4c", false},
		{"Ks", false},
		{"Kx", true},
		{"4s", true},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			asCard, err := ParseCard(tt.args)
			if err != nil {
				if !tt.wantErr {
					t.Error("error not expected", err)
				}
				return
			}
			asString := asCard.String()
			if asString != tt.args {
				t.Errorf("expected: %s got: %s", tt.args, asString)
			}
		})
	}
}
