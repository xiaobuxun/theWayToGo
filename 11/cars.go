package main

type car struct {
	model string
	manuFactory string
	buildYear int
}

type cars []*car

func (cs cars) process(f func(car *car)) {
	for _,c := range cs {
		f(c)
	}
}

func (cs cars) fintAll(f func(car *car) bool) cars {
	cars := make([]*car, 0)

	cs.process(func(c *car) {
		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}

func (cs cars) carMap(f func(car *car) interface{}) []interface{} {
	result := make([]interface{}, 0)
	ix := 0

	cs.process(func (c *car){
		result[ix] = f(c)
		ix++
	})

	return result
}