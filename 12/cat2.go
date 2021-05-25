package main

import (
	"flag"
	"fmt"
	"os"
)

func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		nr, re := f.Read(buf[:])
		switch {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat:error reading:%s\n", re)
			os.Exit(1)
		case nr == 0:
			return
		case nr > 0:
			nw, we := os.Stdout.Write(buf[0:nr])
			if nw != nr {
				fmt.Fprintf(os.Stderr, "cat:write error:%s\n", we.Error())
			}
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	} else {
		for i := 0; i < flag.NArg(); i++ {
			f,err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprintf(os.Stderr, "cat: can't open file %s\n", flag.Arg(i))
				os.Exit(1)
			}
			cat(f)
			f.Close()
		}
	}
}