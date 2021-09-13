package main

import (
  "os"
)

func main() {
  grid := NewGrid(4,4)
  Generate(grid)

  grid.Output(os.Stdout)
}
