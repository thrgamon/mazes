package main

import "math/rand"

type Grid struct {
	rows   int
	cols   int
	matrix CellMatrix
}

type CellRow []*Cell
type CellMatrix []CellRow

func NewGrid(rows int, cols int) *Grid {
	var g Grid

	g.rows = rows
	g.cols = cols
	g.matrix = prepareGrid(rows, cols)
	configureCells(g)

	return &g
}

func (g Grid) GetCell(row int, col int) *Cell {
	var cell *Cell

	if row < 0 || row > g.rows {
		return cell
	}

	if col < 0 || col > g.cols {
		return cell
	}

	cell = g.matrix[row][col]

	return cell
}

func (g Grid) RandomCell() *Cell {
	randomRow := rand.Intn(g.rows)
	randomCol := rand.Intn(g.cols)
	cell := g.GetCell(randomRow, randomCol)

	return cell
}

func (g Grid) Size() int {
	return g.rows * g.cols
}

func (g Grid) EachRow(f func(CellRow)) {
	for r := 0; r < g.rows-1; r++ {
		f(g.matrix[r])
	}
}

func (g Grid) EachCell(f func(*Cell)) {
	for r := 0; r < g.rows-1; r++ {
		for c := 0; c < g.cols-1; c++ {
			cell := g.matrix[r][c]

			if cell != nil {
				f(cell)
			}
		}
	}
}

func prepareGrid(rows int, cols int) CellMatrix {
	grid := make(CellMatrix, rows)

	for r := range grid {
		grid[r] = make(CellRow, cols)

		for c := range grid[r] {
			grid[r][c] = NewCell(r, c)
		}
	}

	return grid
}

func configureCells(g Grid) {
	g.EachCell(func(cell *Cell) {
		r := cell.row
		c := cell.col

		cell.north = g.GetCell(r-1, c)
		cell.south = g.GetCell(r+1, c)
		cell.west = g.GetCell(r, c-1)
		cell.east = g.GetCell(r, c+1)
	})
}
