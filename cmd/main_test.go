package main

import (
	"testing"
)

type TestCaseSortBalls struct {
	Matrix [][]int
	Answer string
}

func TestСanSortBalls(t *testing.T) {

	cases := []TestCaseSortBalls{
		{
			Matrix: [][]int{{1, 2}, {2, 1}},
			Answer: yes,
		},
		{
			Matrix: [][]int{{10, 20, 30}, {1, 1, 1}, {0, 0, 1}},
			Answer: no,
		},
	}

	for caseNum, item := range cases {

		res := canSortBalls(item.Matrix)

		if res != item.Answer {
			t.Errorf("[%d] wrong Answer: got %d, expected %d",
				caseNum, res, item.Answer)
		}

	}

}
