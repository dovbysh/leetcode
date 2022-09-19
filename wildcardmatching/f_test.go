package wildcardmatching

import (
	"fmt"
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "abefcdgiescdfimde",
				p: "ab*cd?i*de",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "aa",
				p: "*",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "cb",
				p: "?a",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "adceb",
				p: "*b",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "adceb",
				p: "*a*b",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("s = \"%s\", p = \"%s\"", tt.args.s, tt.args.p)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
