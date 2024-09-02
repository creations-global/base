package base

import (
	"os"
	"github.com/spf13/cobra"
)

type Base struct {
	flags          Flags
	rootCmd        *cobra.Command
	buildTimestamp string
}

func (what *Base) Execute() {
	// Your code goes here.
}