package leetcode

import "fmt"

func generate(numRows int) [][]int {
	var ret [][]int
	var line []int
	ret = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		line = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j>>1+1 > i-1 {
				line[j] = 1
			} else {
				line[j] = ret[i-1][j>>1] + ret[i-1][j>>1+1]
			}
		}
		ret[i] = line
	}
	return ret
}

func test() {

	aa := generate(4)
	for i := range aa {
		fmt.Println(aa[i])
	}
}
