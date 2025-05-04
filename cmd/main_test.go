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
		{
			Matrix: [][]int{{10, 20}, {20, 10}},
			Answer: yes,
		},
		{
			Matrix: [][]int{{1, 0}, {0, 1}},
			Answer: yes,
		},
		{
			Matrix: [][]int{{1, 1}, {1, 1}},
			Answer: yes,
		},
		{
			Matrix: [][]int{{1, 1}, {0, 0}},
			Answer: no,
		},
	}

	for caseNum, item := range cases {

		res := canSortBalls(item.Matrix)
		if res != item.Answer {
			t.Errorf("[%d] wrong Answer: got %s, expected %s",
				caseNum, res, item.Answer)
		}
	}
}

func TestСanSortBallsWrongSizeMatrix(t *testing.T) {

	cases := []TestCaseSortBalls{
		{
			Matrix: [][]int{{1, 2}, {2, 1}, {0, 0}},
			Answer: no,
		},
		{
			Matrix: [][]int{{10, 20, 0}, {20, 10, 0}},
			Answer: no,
		},
	}

	for caseNum, item := range cases {

		res := canSortBalls(item.Matrix)
		if res != item.Answer {
			t.Errorf("[%d] wrong Answer: got %s, expected %s",
				caseNum, res, item.Answer)
		}
	}
}
