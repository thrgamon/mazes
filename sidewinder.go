package main

import (
  "math/rand"
  "time"
)

func Sidewinder (g *Grid) *Grid {
  rand.Seed(time.Now().UnixNano())

  g.EachRow(func(cr CellRow){
    var run []*Cell

    for _, c := range cr {
      if c == nil {
        continue
      }

      run = append(run, c)

      atEasternBoundary := c.east == nil
      atSouthernBoundary := c.south == nil

      shouldClose := atEasternBoundary || (!atSouthernBoundary && rand.Intn(2) == 0)

      if shouldClose {
        index := rand.Intn(len(run))
        member := run[index]

        if member.south != nil {
          member.Link(member.south)
        }

        run = []*Cell{}
       } else {
        c.Link(c.east)
      }
    }

  })

  return g
}
