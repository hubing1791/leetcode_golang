package medium_01_08_zero_matrix_lcci

//https://leetcode.cn/problems/zero-matrix-lcci/
//2022-07-22

func setZeroes(matrix [][]int) {
	lineNum, lineLength := len(matrix), len(matrix[0])
	var colHasZero = false
	for _, line := range matrix {
		if line[0] == 0 {
			colHasZero = true
		}
		for lineIndex := 1; lineIndex < lineLength; lineIndex += 1 {
			if line[lineIndex] == 0 {
				line[0] = 0
				matrix[0][lineIndex] = 0
			}
		}
	}

	for i := lineNum - 1; i >= 0; i-- {
		// j不可以从0开始，会被提前更新导致把后面的数都置为0
		for j := 0; j < lineLength; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if colHasZero {
			matrix[i][0] = 0
		}
	}

}
