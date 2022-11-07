package t27_27424

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func T2slice(filename string) ([][2]int, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read a %s", err)
	}
	s := strings.Split(string(b), "\n")
	c, err := strconv.Atoi(strings.TrimSpace(s[0]))
	if err != nil {
		return nil, fmt.Errorf("Atoi c %s", err)
	}
	zz := make([][2]int, c)
	for i := 0; i < c; i++ {
		x := strings.Fields(s[i+1])
		x0, err := strconv.Atoi(strings.TrimSpace(x[0]))
		if err != nil {
			return nil, fmt.Errorf("Atoi x0 %s", err)
		}
		x1, err := strconv.Atoi(strings.TrimSpace(x[1]))
		if err != nil {
			return nil, fmt.Errorf("Atoi x1 %s", err)
		}
		zz[i] = [2]int{x0, x1}
	}
	return zz, nil
}
func Test_strageSum(t *testing.T) {
	type args struct {
		a [][2]int
	}
	type tt struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	tests := []tt{
		{
			name: "test 32",
			args: args{
				a: [][2]int{
					{1, 3},
					{5, 12},
					{6, 9},
					{5, 4},
					{3, 3},
					{1, 1},
				},
			},
			want:    32,
			wantErr: false,
		},
	}
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("pwd %s", err)
		return
	}
	zz, err := T2slice(pwd + "/27-A_demo.txt")
	tests = append(tests, tt{
		name:    "27-A_demo.txt",
		args:    args{a: zz},
		want:    0,
		wantErr: false,
	})
	zz, err = T2slice(pwd + "/27-B_demo.txt")
	tests = append(tests, tt{
		name:    "27-B_demo.txt",
		args:    args{a: zz},
		want:    0,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strageSum(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("strageSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("strageSum() got = %v, want %v", got, tt.want)
			}
		})
	}
}
