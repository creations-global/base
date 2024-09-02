package base

import (
	"github.com/spf13/cobra"
)

const ()

func (what *Base) initRoot() *Base {
	what.rootCmd = &cobra.Command{}

	return what
}
