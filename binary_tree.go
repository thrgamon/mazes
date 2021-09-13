package main

import "math/rand"

func Generate (g Grid) Grid {
  g.EachCell(func(c *Cell){
    var neighbours []*Cell
    if c.north != nil {
      neighbours = append(neighbours, c.north)
    }

    if c.east != nil {
      neighbours = append(neighbours, c.east)
    }

    index := rand.Intn(len(neighbours))

    neighbour := neighbours[index]

    if neighbour != nil {
      c.Link(neighbour)
    }
  })

  return g
}
