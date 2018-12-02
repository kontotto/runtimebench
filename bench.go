package main

type Bench interface {
	Run() (*Result, error)
}

type Result struct {
	Create int
	Delete int
	Kill   int
	Start  int
}
