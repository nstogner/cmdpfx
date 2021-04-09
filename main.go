package main

import (
	"fmt"
	"os"
	"os/exec"

	"errors"

	"github.com/nstogner/cmdpfx/prefix"
)

func main() {
	if err := prefix.RunCommand(os.Args[1], os.Args[2:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if ee, ok := errors.Unwrap(err).(*exec.ExitError); ok {
			os.Exit(ee.ExitCode())
		}
		os.Exit(1)
	}
}
