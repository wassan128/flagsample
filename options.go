package main

type Options struct {
	Opt1 string   `flag:"opt1" cfg:"opt1"`
	Opt2 string   `flag:"opt2" cfg:"opt2"`
	Opt3 []string `flag:"opt3" cfg:"opt3"`
	Opt4 []int    `flag:"opt4" cfg:"opt4"`
	Opt5 Struct   `flag:"opt5" cfg:"opt5"`
}

func NewOptions() *Options {
	return &Options{
		Opt1: "opt1-val",
		Opt2: "opt2-val",
		Opt3: []string{"opt3-val"},
		Opt4: []int{0},
		Opt5: Struct{},
	}
}
