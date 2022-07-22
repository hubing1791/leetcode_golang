package medium_189_rotate_array

import "testing"

func Test_rotateTwoPointer(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	var tests = []struct {
		name string
		args args
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateTwoPointerRight(tt.args.nums, tt.args.k)
		})
	}
}
