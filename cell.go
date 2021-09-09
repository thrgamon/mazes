package main

type Cell struct {
  row int
  col int
  links map[*Cell]bool

  north *Cell
  east *Cell
  south *Cell
  west *Cell
}

func NewCell(row int, col int) *Cell {
  var c Cell

  c.links = make(map[*Cell]bool)
  c.row = row
  c.col = col

  return &c
}

func (c Cell) Link(c2 *Cell, bidi bool) Cell {
  c.links[c2] = true
  if bidi {
    c2.Link(&c, false)
  }
  return c
}

func (c Cell) Unlink(c2 *Cell, bidi bool) Cell {
  delete(c.links, c2)
  if bidi {
    c2.Unlink(&c, false)
  }
  return c
}

func (c Cell) Links() []*Cell {
  keys := make([]*Cell, len(c.links))

  i := 0
  for k := range c.links {
      keys[i] = k
      i++
  }

  return keys
}

func (c Cell) Linked(c2 *Cell) bool {
  _, ok := c.links[c2]

  return ok
}

func (c Cell) Neighbours() []*Cell {
  var neighbours []*Cell

  if c.north != nil {
    neighbours = append(neighbours, c.north)
  }

  if c.east != nil {
    neighbours = append(neighbours, c.east)
  }

  if c.south != nil {
    neighbours = append(neighbours, c.south)
  }

  if c.west != nil {
    neighbours = append(neighbours, c.west)
  }

  return neighbours
}
