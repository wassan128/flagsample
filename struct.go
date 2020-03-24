package main

type Struct struct {
	name  string
	value int
}

func (a *Struct) Get() interface{} {
	s := Struct{}
	s.name = a.name
	s.value = a.value
	return s
}

func (a *Struct) Set(s string) error {
	return nil
}

func (a *Struct) String() string {
	return ""
}
