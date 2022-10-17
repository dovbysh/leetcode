package course_schedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0->1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "0->1, 1->0",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "0->1, 1->2",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 0}, {2, 1}},
			},
			want: true,
		},
		{
			name: "0->1, 2->1, 1->0",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 0}, {1, 2}, {0, 1}},
			},
			want: false,
		},
		{
			name: "0->1, 1->2, 3",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 1}},
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				numCourses:    1,
				prerequisites: nil,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				numCourses:    2,
				prerequisites: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
