package main

import (
	"fmt"
	"os/exec"
	"time"
)

type dummyBench struct {
	Result
}

func (b *dummyBench) Run() (*Result, error) {
	_, err := b.Create()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *dummyBench) Create() (int, error) {
	command := exec.Command("echo", "create")
	before := time.Now()
	err := command.Run()
	if err != nil {
		return 0, err
	}
	after := time.Now()
	fmt.Printf("%+v", after.Nanosecond()-before.Nanosecond())
	return 0, nil
}

func (b *dummyBench) Delete() (int, error) {
	return 0, nil
}

func (b *dummyBench) Kill() (int, error) {
	return 0, nil
}

func (b *dummyBench) Start() (int, error) {
	return 0, nil
}
