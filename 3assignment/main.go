package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
