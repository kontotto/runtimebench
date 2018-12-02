package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDummyFactoryCreate(t *testing.T) {
	factory := &dummyFactory{}

	want := &dummyBench{}
	got, err := factory.Create()

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
