package base

import (
	"github.com/spf13/cobra"
)

const ()

func (what *Base) initRoot() *Base {
	what.rootCmd = &cobra.Command{
		Use:           "base",
		Version:       BaseVersion,
		Short:         "\n" + Logo,
	}

	return what
}
