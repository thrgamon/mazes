package main

import (
  "math/rand"
  "time"
)

func BinaryTree (g *Grid) *Grid {
  rand.Seed(time.Now().UnixNano())

  g.EachCell(func(c *Cell){
    var neighbours []*Cell

    // The book implements this going to the north
    // but it doesn't work unless we go south.
    // I am not really sure why...
    if c.south != nil {
      neighbours = append(neighbours, c.south)
    }

    if c.east != nil {
      neighbours = append(neighbours, c.east)
    }

    if len(neighbours) != 0 {
      index := rand.Intn(len(neighbours))

      neighbour := neighbours[index]

      if neighbour != nil {
        c.Link(neighbour)
      }
    }
  })

  return g
}
