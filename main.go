package main

import (
  "os"
)

func main() {
  grid := NewGrid(10,10)
  Sidewinder(grid)

  grid.Output(os.Stdout)
}
