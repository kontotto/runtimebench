package main

type dummyBench struct {
	Result
}

func (b *dummyBench) Run() (*Result, error) {
	return nil, nil
}
