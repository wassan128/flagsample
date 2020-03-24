package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	options "github.com/mreiferson/go-options"
)

type config map[string]interface{}

func main() {
	flagSet := flag.NewFlagSet("flagsample", flag.ExitOnError)

	opt3val := StringArray{}
	opt4val := IntArray{}
	opt5val := Struct{}

	flagSet.String("opt1", "default-value-of-opt1", "Description of opt1")
	flagSet.String("opt2", "default-value-of-opt2", "Description of opt2")
	flagSet.Var(&opt3val, "opt3", "Description of opt3")
	flagSet.Var(&opt4val, "opt4", "Description of opt4")
	flagSet.Var(&opt5val, "opt5", "Description of opt5")

	flagSet.Parse(os.Args[1:])
	fmt.Printf("flagSet: %+v\n", flagSet)

	var cfg config
	configFile := flagSet.Lookup("opt1").Value.String()
	if configFile != "" {
		_, err := toml.DecodeFile(configFile, &cfg)
		if err != nil {
			fmt.Printf("Failed to load config file")
		}
	}

	opts := NewOptions()
	options.Resolve(opts, flagSet, cfg)

	fmt.Printf("cfg: %#v\n", cfg)
	fmt.Printf("opts: %#v\n", opts)
}
