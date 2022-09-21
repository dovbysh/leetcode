package reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "123",
			x:    123,
			want: 321,
		},
		{
			name: "-123",
			x:    -123,
			want: -321,
		},
		{
			name: "120",
			x:    120,
			want: 21,
		},
		{
			name: "1534236469",
			x:    1534236469,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
