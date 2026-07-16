package main

import (
	"fmt"
	"os"
)

func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat;errorreading %s\n", err.Error())
			os.Exit(1)
		case nr == 0: //eof
			return
		case nr > 0:
			os.Stdout.Write(buf[:nr])
		}
	}
}

func main() {

}
