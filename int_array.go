package main

import "strconv"

type IntArray []int

func (a *IntArray) Get() interface{} {
	return []int(*a)
}

func (a *IntArray) Set(s string) error {
	i, _ := strconv.Atoi(s)
	*a = append(*a, i)
	return nil
}

func (a *IntArray) String() string {
	return ""
}
