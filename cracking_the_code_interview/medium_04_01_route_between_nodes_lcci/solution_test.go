package medium_04_01_route_between_nodes_lcci

import "testing"

func Test_findWhetherExistsPathBFS(t *testing.T) {
	type args struct {
		n      int
		graph  [][]int
		start  int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case one", args{3, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}, 0, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWhetherExistsPathBFS(tt.args.n, tt.args.graph, tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("findWhetherExistsPathBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
