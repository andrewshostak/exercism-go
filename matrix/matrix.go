package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(input string) (Matrix, error) {
	rowsString := strings.Split(input, "\n")
	matrixString := make([][]string, len(rowsString), len(rowsString))
	for i := range rowsString {
		trimmed := strings.TrimSpace(rowsString[i])
		matrixString[i] = strings.Split(trimmed, " ")
		if i == 0 {
			continue
		}
		if len(matrixString[i]) != len(matrixString[i-1]) {
			return nil, errors.New("bad input")
		}
	}

	matrix := make([][]int, 0, len(matrixString))
	for i := range matrixString {
		row := make([]int, 0, len(matrixString[i]))
		for j := range matrixString[i] {
			intItem, err := strconv.ParseInt(matrixString[i][j], 10, 64)
			if err != nil {
				return nil, err
			}
			row = append(row, int(intItem))
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}

func (m *Matrix) Rows() [][]int {
	var matrix [][]int = *m
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func (m *Matrix) Cols() [][]int {
	var matrix [][]int = *m
	cols := make([][]int, 0, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		col := make([]int, 0, len(matrix))
		for j := range matrix {
			col = append(col, matrix[j][i])
		}
		cols = append(cols, col)
	}

	return cols
}

func (m *Matrix) Set(r, c, value int) bool {
	if r < 0 || c < 0 {
		return false
	}

	if r > (len(*m)-1) || c > (len(*m)-1) {
		return false
	}

	matrix := m.Rows()
	matrix[r][c] = value
	*m = matrix

	return true
}
