package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		opt1 = flag.String("opt1", "default-value-of-opt1", "Description of opt1")
		opt2 = flag.String("opt2", "default-value-of-opt2", "Description of opt2")
	)

	flag.Parse()
	fmt.Printf("opt1: %+v\n", *opt1)
	fmt.Printf("opt2: %+v\n", *opt2)
	fmt.Printf("args: %+v\n", flag.Args())
}
