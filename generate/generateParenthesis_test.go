package generate

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n = 1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n = 2",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "n = 3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n = 4",
			n:    4,
			want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = \n%v,\n want \n%v", got, tt.want)
			}
		})
	}
}
