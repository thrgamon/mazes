package main

type Grid struct {
  rows int
  cols int
  matrix CellMatrix
}

type CellMatrix [][]*Cell

func NewGrid(rows int, cols int) *Grid {
  var g Grid

  g.rows = rows
  g.cols = cols
  g.matrix = prepareGrid(rows, cols)
  configureCells(g)

  return &g
}

func (g Grid) Coord (row int, col int) *Cell{
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

func prepareGrid(rows int, cols int) CellMatrix {
  grid := make(CellMatrix, rows)

  for r := range grid {
    grid[r] = make([]*Cell, cols)

    for c := range grid[r] {
      grid[r][c] = NewCell(r, c)
    }
  }

  return grid
}

func configureCells(g Grid) {
  for r := 0; r < g.rows - 1; r++ {
    for c := 0; c < g.cols - 1; c++ {
      cell := g.matrix[r][c] 

      cell.north = g.Coord(r - 1,c)
      cell.south = g.Coord(r + 1,c)
      cell.west = g.Coord(r,c - 1)
      cell.east = g.Coord(r,c + 1)
    }
  }
}

