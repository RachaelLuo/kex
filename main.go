package main

import (
	"os"

	"github.com/RachaelLuo/kex/cmd"

	"k8s.io/component-base/cli"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	command := cmd.NewCommand("kex", os.Stdin, os.Stdout, os.Stderr)
	if err := cli.RunNoErrOutput(command); err != nil {
		// Pretty-print the error and exit with an error.
		cmdutil.CheckErr(err)
	}
}
