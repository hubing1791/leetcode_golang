package medium_02_04_parttition_list_lcci

import (
	ll "leetcode_golang/custom_moudle/singly_linked_list"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ll.ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ll.ListNode
	}{
		{"one", args{ll.Generator([]int{1, 4, 3, 2, 5, 2}), 3}, ll.Generator([]int{1, 2, 2, 4, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
