package medium_04_01_route_between_nodes_lcci

//https://leetcode.cn/problems/route-between-nodes-lcci/
//2022-11-20

//别人写好的版本，写的比较漂亮模仿了一下
func findWhetherExistsPathBFS(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int, n)
	Flag := make([]int, n)
	for i := 0; i < len(graph); i++ {
		edges[graph[i][0]] = append(edges[graph[i][0]], graph[i][1])
	}
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		Flag[tmp] = 1
		if tmp == target {
			return true
		}
		for i := 0; i < len(edges[tmp]); i++ {
			if Flag[edges[tmp][i]] == 0 {
				queue = append(queue, edges[tmp][i])
			}
		}
	}
	return false
}

//进行一个仿写
func findWhetherExistsPathBFSMy(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int, n)
	flag := make([]int, n)
	for i := 0; i < len(graph); i++ {
		edges[graph[i][0]] = append(edges[graph[i][0]], graph[i][1])
	}

	queue := make([]int, 0)
	queue = append(queue, start)
	index := 0
	for len(queue) > index {
		tmp := queue[index]
		index += 1
		if tmp == target {
			return true
		}
		flag[tmp] = 1
		for i := 0; i < len(edges[tmp]); i++ {
			if flag[edges[tmp][i]] == 0 {
				queue = append(queue, edges[tmp][i])
			}
		}
	}
	return false
}

func findWhetherExistsPathDFSMy(n int, graph [][]int, start int, target int) bool {

	// visited := make([]bool, n) 这么写不对，长度不一定为n。
	visited := make([]bool, len(graph))

	return DFSFind(graph, visited, start, target)
}

func DFSFind(graph [][]int, visited []bool, start, target int) bool {
	// 首先确定没有访问过。
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			if graph[i][0] == start && graph[i][1] == target {
				return true
			}
			visited[i] = true
			// 这个逆向搜索很重要，可以防止搜索空间爆炸
			if graph[i][1] == target && DFSFind(graph, visited, start, graph[i][0]) {
				return true
			}
			visited[i] = false
		}
	}
	return false

}
