package main

import "fmt"

type op func(a,b int) int

type request struct {
	a, b int
	reply chan int
}

func run(op op, req *request) {
	req.reply <- op(req.a, req.b)
}

func server (op op, requests chan *request, quit chan bool) {
	for {
		select {
		case req := <- requests:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op op) (requests chan *request, quit chan bool) {
	requests = make(chan *request)
	quit = make(chan bool)
	go server(op, requests, quit)
	return requests, quit
}

func main() {
	requests, quit := startServer(func(a, b int) int {
		return a + b
	})

	var reqs [2]request

	for i:=0;i<len(reqs);i++ {
		req := &reqs[i]
		req.a = i
		req.b = i * 2
		req.reply = make(chan int)
		requests <- req
	}

	for i:=0;i<len(reqs);i++ {
		fmt.Printf("reply in %d is:%d\n", i, <-reqs[i].reply)
	}
	quit <- true
	fmt.Println("done")
}