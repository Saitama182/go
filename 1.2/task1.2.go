package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", ind+1, arg)
	}
}
