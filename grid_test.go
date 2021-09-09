package main

import(
  "testing"
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
