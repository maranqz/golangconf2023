package main

import (
	"fmt"

	"ddd/testdata/src/write/nested"
	"ddd/testdata/src/write/nested/nested2"
)

func main() {
	n := nested.Nested{}
	n.Write = 2  // want `nested.Nested.Write is readonly`
	n.Write += 2 // want `nested.Nested.Write is readonly`
	n.Write++    // want `nested.Nested.Write is readonly`
	n.Write -= 2 // want `nested.Nested.Write is readonly`
	n.Write--    // want `nested.Nested.Write is readonly`
	_ = n.Read
	fmt.Println(n.Read)

	n2 := nested2.Nested{}
	n2.Write = "2"  // want `nested2.Nested.Write is readonly`
	n2.Write += "2" // want `nested2.Nested.Write is readonly`
	_ = n2.Read
	fmt.Println(n2.Read)
}
