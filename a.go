package main

import (
	"fmt"
	"os"
	"time"
)

type A interface {
	AA()
	BB()
}

type B struct {
	l int
}

func (b *B) AA() {
	fmt.Print("1")
}

func (b *B) BB() {
	fmt.Print("2")
}

var _ A = (*B)(nil)

func main() {

	var a time.Time
	var b time.Duration

	a = time.Now()
	b = time.Since(a)

	fmt.Println(a, "aaaa", b)

	c, err := os.Open("ss")

	d, err := c.Stat()

	fmt.Print(d, err)

}
