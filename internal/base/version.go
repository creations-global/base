package base

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (what *Base) initVersion() *Base {
	what.rootCmd.AddCommand(&cobra.Command{
		Use:   PrintVersionCommand,
		Short: "Get version information",
		Long:  "\n" + Logo + "\n\n" + fmt.Sprintf(VersionText, what.buildTimestamp),
	})

	return what
}