// package name: reef
package main

import (
	"C"
	"fmt"

	"github.com/exis-io/riffle"
)

//export Tester
func Tester() string {
	fmt.Println("Starting")
	return riffle.Tester("asef")
}

func main() {}
