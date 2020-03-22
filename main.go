package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagSet := flag.NewFlagSet("flagsample", flag.ExitOnError)

	opt3val := StringArray{}
	opt4val := IntArray{}

	flagSet.String("opt1", "default-value-of-opt1", "Description of opt1")
	flagSet.String("opt2", "default-value-of-opt2", "Description of opt2")
	flagSet.Var(&opt3val, "opt3", "Description of opt3")
	flagSet.Var(&opt4val, "opt4", "Description of opt4")

	flagSet.Parse(os.Args[1:])
	fmt.Printf("%+v", flagSet)
}
