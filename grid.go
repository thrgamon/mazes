package main

import (
  "math/rand"
  "io"
  "strings"
  "fmt"
)

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

	if row < 0 || row > g.rows - 1 {
		return cell
	}

	if col < 0 || col > g.cols - 1 {
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
	for r := 0; r < g.rows; r++ {
		f(g.matrix[r])
	}
}

func (g Grid) EachCell(f func(*Cell)) {
  g.EachRow(func(cr CellRow) {
    for _, cell := range cr {

			if cell != nil {
				f(cell)
			}
    }
  })
}

func (g Grid) Output(writer io.Writer) {
  var output strings.Builder

  fmt.Fprintf(&output, "+%s\n", strings.Repeat("---+", g.cols))

  g.EachRow(func(cr CellRow){
    var top strings.Builder
    fmt.Fprint(&top, "|")

    var bottom strings.Builder
    fmt.Fprint(&bottom, "+")

    for _, cell := range cr {
      if cell == nil {
        cell = NewCell(-1, -1)
      }

      fmt.Fprint(&top, "   ")

      if cell.Linked(cell.east) {
        fmt.Fprint(&top, " ")
      } else {
        fmt.Fprint(&top, "|")
      }

      if cell.Linked(cell.south) {
        fmt.Fprint(&bottom, "   ")
      } else {
        fmt.Fprint(&bottom, "---")
      }

      fmt.Fprint(&bottom, "+")
    }

      fmt.Fprint(&output, top.String())
      fmt.Fprint(&output, "\n")
      fmt.Fprint(&output, bottom.String())
      fmt.Fprint(&output, "\n")

  })

  fmt.Fprintf(writer, "%s", output.String())
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
