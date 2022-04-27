package main

import (
	"fmt"
	"os"
)

// Print a list of arguments to Standard Error Stream to trace the execution of the program
func Trace(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
