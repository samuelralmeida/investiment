package model

import "testing"

func TestAtivos_IsValid(t *testing.T) {
	tests := []struct {
		name string
		a    *Ativos
		want bool
	}{
		{
			name: "nil",
			a:    nil,
			want: false,
		},
		{
			name: "empty",
			a:    &Ativos{},
			want: false,
		},
		{
			name: "success",
			a:    &Ativos{{ID: 1}, {ID: 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsValid(); got != tt.want {
				t.Errorf("Ativos.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
