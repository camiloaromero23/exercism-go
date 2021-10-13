package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(in string) (*Matrix, error) {
	cols := strings.Split(in, "\n")
	var m Matrix = make(Matrix, len(cols))
	for i, c := range cols {
		c = strings.TrimSpace(c)
		r := strings.Split(c, " ")
		if len(r) == 0 {
			return nil, fmt.Errorf("Empty row")
		}
		for j, v := range r {
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			if m[i] == nil {
				m[i] = make([]int, len(r))
			}
			if i > 0 && len(m[i-1]) != len(m[i]) {
				return nil, fmt.Errorf("Uneven rows")
			}
			m[i][j] = val
		}
	}
	return &m, nil
}

func (m Matrix) Rows() [][]int {
	rows := make(Matrix, len(m))
	for i, v := range m {
		for j, r := range v {
			if rows[i] == nil {
				rows[i] = make([]int, len(m[i]))
			}
			rows[i][j] = r
		}
	}
	return rows
}
func (m Matrix) Cols() [][]int {
	cols := make(Matrix, len(m[0]))
	for i, v := range m {
		for j, c := range v {
			if cols[j] == nil {
				cols[j] = make([]int, len(m))
			}
			cols[j][i] = c
		}
	}
	return cols
}
func (m *Matrix) Set(r, c, val int) bool {
	matrix := *m
	if r >= len(matrix[0]) || r < 0 {
		return false
	}
	if c >= len(matrix) || c < 0 {
		return false
	}
	matrix[r][c] = val
	return true
}
