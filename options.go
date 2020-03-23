package main

type Options struct {
	Opt1 string `flag:"opt1" cfg:"opt1"`
	Opt2 string `flag:"opt2" cfg:"opt2"`
	Opt3 string `flag:"opt3" cfg:"opt3"`
}

func NewOptions() *Options {
	return &Options{
		Opt1: "opt1-val",
		Opt2: "opt2-val",
		Opt3: "opt3-val",
	}
}
