package main

type dummyFactory struct {
}

func (f *dummyFactory) Create() (Bench, error) {
	return &dummyBench{}, nil
}
