package main

import(
  "testing"
  "reflect"
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
