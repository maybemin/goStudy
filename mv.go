package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	err := os.Rename(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
}
