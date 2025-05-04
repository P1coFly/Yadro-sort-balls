package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const (
	yes = "yes"
	no  = "no"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	fmt.Fprint(out, canSortBalls(matrix))
}

func canSortBalls(matrix [][]int) string {
	if len(matrix[0]) != len(matrix) {
		return no
	}

	rowSums := make([]int, len(matrix))
	columnSums := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			rowSums[i] += matrix[i][j]
			columnSums[j] += matrix[i][j]
		}
	}

	sortDescending(columnSums)
	sortDescending(rowSums)

	for i := 0; i < len(matrix); i++ {
		if columnSums[i] != rowSums[i] {
			return no
		}
	}
	return yes
}

func sortDescending(numbers []int) {
	slices.SortFunc(numbers, func(a, b int) int {
		return b - a
	})
}
