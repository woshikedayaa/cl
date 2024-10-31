package in

import (
	"github.com/spf13/cobra"
	"github.com/woshikedayaa/cl/cmd/cl"
	"golang.design/x/clipboard"
	"io"
	"os"
	"time"
)

var inCommand = &cobra.Command{
	Use:   "in",
	Short: "write to the clipboard",
	Long:  "write to the clipboard",
	RunE:  In,
}

var (
	flagSource string
	flagQuiet  bool
	flagImage  bool
)

func init() {
	inCommand.Flags().StringVarP(&flagSource, "source", "s", "stdin", "source to write the clipboard")
	inCommand.Flags().BoolVarP(&flagQuiet, "quiet", "q", false, "disable output to the stdout")
	inCommand.Flags().BoolVarP(&flagImage, "image", "i", false, "specifies that the source content is of type Image")

	cl.RootCmd.RunE = inCommand.RunE
	cl.RootCmd.Run = inCommand.Run
	cl.RootCmd.AddCommand(inCommand)
}

func In(cmd *cobra.Command, args []string) error {
	f := clipboard.FmtText
	if flagImage {
		f = clipboard.FmtImage
	}
	reader := io.Reader(nil)
	switch flagSource {
	case "stdin":
		reader = os.Stdin
	default:
		f, err := os.Open(flagSource)
		if err != nil {
			return err
		}
		defer f.Close()
		reader = f
	}
	cnt, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	if !flagQuiet {
		cmd.Println(string(cnt))
	}
	// FIXME
	// Issues: https://github.com/golang-design/clipboard/issues/66
	_ = clipboard.Write(f, cnt)
	time.Sleep(time.Second)
	return nil
}
