package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	salary float64
	chF chan func()
}

func (p *person) backend() {
	for f := range p.chF {
		f()
	}
}

func (p *person) setSalary(s float64) {
	p.chF <- func() {
		p.salary = s
	}
}

func (p *person) getSalary() float64 {
	fch := make(chan float64)
	p.chF <- func() {
		fch <- p.salary
	}

	return <- fch
}

func (p *person) String() string {
	return "Person - name is:" + p.name + " salary is:" + strconv.FormatFloat(p.getSalary(), 'f', 2, 64)
}

func newPerson(name string, salary float64) *person {
	p := &person{name, salary, make(chan func())}

	go p.backend()

	return p
}

func main() {
	bs := newPerson("li", 4000.25)
	fmt.Println(bs)
	bs.setSalary(5000.89)
	fmt.Println("salary changed:")
	fmt.Println(bs)
}