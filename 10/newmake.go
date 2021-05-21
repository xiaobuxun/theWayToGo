package main

type one map[string]string
type two struct {
	attr1 string
	attr2 int
}

func main() {
	x := make(one)
	x["a"] = "a"
	x["b"] = "b"
	y := new(two)
	y.attr1 = "A"
	y.attr2 = 1
}