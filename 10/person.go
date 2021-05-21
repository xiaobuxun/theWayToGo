package main

import "strings"

type person struct {
	firstName string
	lastName string
}

func upperPerson(p *person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

