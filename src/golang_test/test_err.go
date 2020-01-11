package main

import "fmt"

type ErrMine struct {
	Err string
}

func (p *ErrMine) Error() string {
	return p.Err
}

func main() {
	a := sometest()
	fmt.Println(a)
}

func sometest() error {
	e := new(ErrMine)
	e.Err = "jii"
	return e
}
