package linkedlist_test

import (
	"fmt"

	"github.com/devlights/linkedlist"
)

func ExampleNewCircular() {
	var (
		circular = linkedlist.NewCircular[int](3)
	)
	fmt.Println(circular)

	// Output:
	// []
}

func ExampleCircular_Add() {
	var (
		circular = linkedlist.NewCircular[string](5)
	)

	for _, r := range "helloworld" {
		circular.Add(string(r))
		fmt.Println(circular)
	}

	// Output:
	// [h ->]
	// [h -> e ->]
	// [h -> e -> l ->]
	// [h -> e -> l -> l ->]
	// [h -> e -> l -> l -> o ->]
	// [e -> l -> l -> o -> w ->]
	// [l -> l -> o -> w -> o ->]
	// [l -> o -> w -> o -> r ->]
	// [o -> w -> o -> r -> l ->]
	// [w -> o -> r -> l -> d ->]
}

func ExampleCircular_ToSlice() {
	var (
		circular = linkedlist.NewCircular[string](5)
	)

	for _, r := range "helloworld" {
		circular.Add(string(r))
	}

	for i, v := range circular.ToSlice() {
		fmt.Printf("%02d: %v\n", i, v)
	}

	// Output:
	// 00: w
	// 01: o
	// 02: r
	// 03: l
	// 04: d
}

func ExampleCircular_Iterate() {
	var (
		circular = linkedlist.NewCircular[string](5)
	)

	for _, r := range "helloworld" {
		circular.Add(string(r))
	}

	if circular.Head != nil {
		for n := circular.Head; ; n = n.Next {
			fmt.Println(n)

			if n == circular.Tail {
				break
			}
		}
	}

	// Output:
	// (v:w,n:o)
	// (v:o,n:r)
	// (v:r,n:l)
	// (v:l,n:d)
	// (v:d,n:nil)
}

func ExampleCircular_Delete() {
	var (
		circular = linkedlist.NewCircular[string](5)
	)

	for _, r := range "helloworld" {
		circular.Add(string(r))
	}

	var (
		fn = func(v1, v2 string) bool {
			return v1 == v2
		}
		deletes = []string{"d", "w", "r"}
	)
	for _, d := range deletes {
		if ok := circular.Delete(d, fn); !ok {
			panic(fmt.Errorf("Delete() returns false (%v)", d))
		}

		fmt.Println(circular)
	}

	// Output:
	// [w -> o -> r -> l ->]
	// [o -> r -> l ->]
	// [o -> l ->]
}
