package main

import (
	"fmt"

	"ddd/testdata/src/write/nested"
	"ddd/testdata/src/write/nested/nested2"
	"ddd/testdata/src/write/writable"
)

type Int struct {
	Write int
	Read  int
}

type String struct {
	Write string
	Read  string
}

func main() {
	n := nested.Nested{}
	n.Write = 2  // want `n.Write is readonly`
	n.Write += 2 // want `n.Write is readonly`
	n.Write++    // want `n.Write is readonly`
	n.Write -= 2 // want `n.Write is readonly`
	n.Write--    // want `n.Write is readonly`
	_ = n.Read
	fmt.Println(n.Read)

	n2 := nested2.Nested{}
	n2.Write = "2"  // want `n2.Write is readonly`
	n2.Write += "2" // want `n2.Write is readonly`
	_ = n2.Read
	fmt.Println(n2.Read)

	// Local

	wInt := Int{}
	wInt.Write = 2
	wInt.Write += 2
	wInt.Write++
	wInt.Write -= 2
	wInt.Write--

	wString := String{}
	wString.Write = "2"
	wString.Write += "2"
	_ = wString.Read
	fmt.Println(wString.Read)

	// External
	weInt := writable.Int{}
	weInt.Write = 2
	weInt.Write += 2
	weInt.Write++
	weInt.Write -= 2
	weInt.Write--

	weString := writable.String{}
	weString.Write = "2"
	weString.Write += "2"
	_ = weString.Read
	fmt.Println(weString.Read)
}
