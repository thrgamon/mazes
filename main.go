package main

import (
	"log"
	"github.com/ajstarks/svgo"
	"net/http"
)

func main() {
	http.Handle("/sidewinder", http.HandlerFunc(sidewinder))
	http.Handle("/binary", http.HandlerFunc(binaryTree))
  http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func sidewinder(w http.ResponseWriter, req *http.Request) {
  grid := NewGrid(20,20)
  Sidewinder(grid)

  renderSvg(grid, w)
}

func binaryTree(w http.ResponseWriter, req *http.Request) {
  grid := NewGrid(20,20)
  BinaryTree(grid)

  renderSvg(grid, w)
}

func renderSvg(grid *Grid, w http.ResponseWriter) {
  cellSize := 30
  offset := 5
  height := grid.rows * cellSize 
  width := grid.cols * cellSize 
  strokeStyle := "stroke:black; stroke-width: 2;"

  w.Header().Set("Content-Type", "image/svg+xml")
  s := svg.New(w)
  s.Start(width + 10, height + 10)

  grid.EachCell(func(c *Cell) {
    x1 := (c.col * cellSize) + offset
    x2 := ((c.col + 1) * cellSize) + offset
    y1 := (c.row * cellSize) + offset
    y2 := ((c.row + 1) * cellSize) + offset


    if c.north == nil {
      s.Line(x1, y1, x2, y1, strokeStyle)
    }

    if c.west == nil {
      s.Line(x1, y1, x1, y2, strokeStyle)
    }

    if !c.Linked(c.east) {
      s.Line(x2, y1, x2, y2, strokeStyle)
    }

    if !c.Linked(c.south) {
      s.Line(x1, y2, x2, y2, strokeStyle)
    }

  })
  // grid.Output(w)
  // s.Line(10, 10, 250, 250, strokeStyle)
  s.End()
}
