package version

import (
	"github.com/spf13/cobra"
	"github.com/woshikedayaa/cl/cmd/cl"
	"runtime"
)

const fromBuildString = "FROM BUILD"

var (
	Version    = fromBuildString
	Branch     = fromBuildString
	CommitHash = fromBuildString
	GoVersion  = runtime.Version()
	Arch       = runtime.GOARCH
	OS         = runtime.GOOS
)

var (
	flagPrintVerbose = false
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version of cl",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("cl version:%s,running on: %s_%s\n", Version, OS, Arch)
		if flagPrintVerbose {
			cmd.Printf("branch: %s,commit: %s,GoVersion: %s\n", Branch, CommitHash, GoVersion)
		}
	},
}

func init() {
	versionCommand.Flags().BoolVarP(&flagPrintVerbose, "verbose", "v", false, "print out build information")

	cl.RootCmd.AddCommand(versionCommand)
}
