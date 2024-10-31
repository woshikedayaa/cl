package out

import (
	"github.com/spf13/cobra"
	"github.com/woshikedayaa/cl/cmd/cl"
)

var outCommand = &cobra.Command{
	Use:   "out",
	Short: "output the current clipboard content",
	RunE:  Out,
}

func Out(cmd *cobra.Command, args []string) error {
	return nil // todo
}

func init() {
	cl.RootCmd.AddCommand(outCommand)
}
