package cl

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "cl",
	Short: "clipboard cli helper",
	Run:   nil, // value fom the cmd/cl/in package
}

func Main() {
	// init clipboard
	err := clipboard.Init()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
