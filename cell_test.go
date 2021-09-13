package main

import(
  "testing"
  "reflect"
)

func TestNewCell(t *testing.T) {
  row := 20
  col := 10
  c1 := NewCell(row, col)

  t.Run("assign rows", func(t *testing.T) {
    got := c1.row
    want := row

    if got != want {
      t.Errorf("got %v want %v given, %v", got, want, c1)
    }
  })

  t.Run("assign cols", func(t *testing.T) {
    got := c1.col
    want := col

    if got != want {
      t.Errorf("got %v want %v given, %v", got, want, c1)
    }
  })
}

func TestLink(t *testing.T) {
  c1 := NewCell(10, 10)
  c2 := NewCell(10, 10)

  c1.Link(c2)
  
  want := c2
  got := c1.links[c2]

  if got != true {
    t.Errorf("got %v want %v given, %v", got, want, c1.links)
  }
}

func TestBLink(t *testing.T) {
  c1 := NewCell(10, 10)
  c2 := NewCell(10, 10)

  c1.BLink(c2, true)
  
  want := c2
  got := c1.links[c2]

  if got != true {
    t.Errorf("got %v want %v given, %v", got, want, c1.links)
  }
}

func TestUnlink(t *testing.T) {
  c1 := NewCell(10, 10)
  c2 := NewCell(10, 10)

  c1.Link(c2)
  c1.Unlink(c2, true)
  
  want := false
  got := c1.links[c2]

  if got != want {
    t.Errorf("got %v want %v given, %v", got, want, c1.links)
  }
}

func TestLinked(t *testing.T) {
  c1 := NewCell(10, 10)
  c2 := NewCell(10, 10)

  c1.Link(c2)
  
  want := true
  got := c1.Linked(c2)

  if got != want {
    t.Errorf("got %v want %v given, %v", got, want, c1.links)
  }
}

func TestLinks(t *testing.T) {
  c1 := NewCell(10, 10)
  c2 := NewCell(10, 10)
  c3 := NewCell(10, 10)

  c1.Link(c2)
  c1.Link(c3)
  
  want := []*Cell{c2, c3}
  got := c1.Links()

  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestNeighbours(t *testing.T) {
  cn := NewCell(10, 10)
  ce := NewCell(10, 10)
  cs := NewCell(10, 10)
  cw := NewCell(10, 10)


  t.Run("full neighbours", func(t *testing.T) {
    c1 := NewCell(10, 10)
    c1.north = cn
    c1.east = ce
    c1.south = cs
    c1.west = cw

    got := c1.Neighbours()
    want := []*Cell{cn, ce, cs, cw}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("half neighbours", func(t *testing.T) {
    c1 := NewCell(10, 10)
    c1.north = cn
    c1.east = ce

    got := c1.Neighbours()
    want := []*Cell{cn, ce}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })
}
