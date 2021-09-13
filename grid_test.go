package main

import(
  "testing"
  "reflect"
  "bytes"
)

func TestNewGrid(t *testing.T) {
  rows := 20
  cols := 10
  g1 := NewGrid(rows, cols)

  t.Run("assign rows", func(t *testing.T) {
    got := g1.rows
    want := rows

    if got != want {
      t.Errorf("got %v want %v given, %v", got, want, g1)
    }
  })

  t.Run("assign cols", func(t *testing.T) {
    got := g1.cols
    want := cols

    if got != want {
      t.Errorf("got %v want %v given, %v", got, want, g1)
    }
  })

  t.Run("length rows", func(t *testing.T) {
    got := len(g1.matrix)
    want := rows

    if got != want {
      t.Errorf("got %v want %v given, %v", got, want, g1)
    }
  })

  t.Run("length cols", func(t *testing.T) {
    got := len(g1.matrix[0])
    want := cols

    if got != want {
      t.Errorf("got %v want %v given, %v", got, want, g1)
    }
  })
}

func TestRandomCell(t *testing.T) {
  rows := 20
  cols := 10
  g1 := NewGrid(rows, cols)

  got := g1.RandomCell()
  var want *Cell

  if reflect.TypeOf(got).Kind() != reflect.TypeOf(want).Kind() {
    t.Errorf("got %v want %v given", got, want)
  }
}

func TestSize(t *testing.T) {
  rows := 20
  cols := 10
  g1 := NewGrid(rows, cols)

  got := g1.Size()
  want := rows * cols

  if got != want {
    t.Errorf("got %v want %v given", got, want)
  }
}

func TestEachCell(t *testing.T) {
  rows := 20
  cols := 10
  g1 := NewGrid(rows, cols)
  counter := 0

  g1.EachCell(func(*Cell) {
    counter += 1
  })
  got := counter
  want := rows * cols

  if got != want {
    t.Errorf("got %v want %v given", got, want)
  }
}

func TestOutput(t *testing.T) {
  rows := 4
  cols := 4

  t.Run("general output", func(t *testing.T) {
    buffer := bytes.Buffer{}
    g1 := NewGrid(rows, cols)
    g1.Output(&buffer)

    want := `+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+`
    got := buffer.String()

    if got != want {
      t.Errorf("got \n%vwant \n%v", got, want)
    }
  })

  t.Run("general output", func(t *testing.T) {
    buffer := bytes.Buffer{}
    g1 := NewGrid(rows, cols)
    c1 := g1.GetCell(1,1)
    c2 := g1.GetCell(2,1)
    c3 := g1.GetCell(1,2)

    c1.Link(c2)
    c1.Link(c3)
    g1.Output(&buffer)

    want := `+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |       |   |
+---+   +---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+`
    got := buffer.String()

    if got != want {
      t.Errorf("got \n%vwant \n%v", got, want)
    }
  })
}
