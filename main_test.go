package main

import (
    "fmt"
    "time"
  "testing"
)

func TestThis(t *testing.T) {
  fmt.Println("hello there")
  t := time.Now()
    y, mon, d := t.Date()
    h, m, s := t.Clock()
    fmt.Println("Year: ", y)
    fmt.Println("Month: ", mon)
    fmt.Println("Day: ", d)
    fmt.Println("Hour: ", h)
    fmt.Println("Minute: ", m)
    fmt.Println("Second: ", s)
}
