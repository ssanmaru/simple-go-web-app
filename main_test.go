package main

import (
    "fmt"
    "time"
  "testing"
)

func TestThis(t *testing.T) {
    fmt.Println("Say Hello there concourse, iam from Go")
    t.Log("Say bye")
  p := time.Now()
    y, mon, d := p.Date()
    h, m, s := p.Clock()
    fmt.Println("Year: ", y)
    fmt.Println("Month: ", mon)
    fmt.Println("Day: ", d)
    fmt.Println("Hour: ", h)
    fmt.Println("Minute: ", m)
    fmt.Println("Second: ", s)
}
