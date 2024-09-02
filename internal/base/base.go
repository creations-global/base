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
	err := what.rootCmd.Execute()
	if err != nil {
		what.rootCmd.Println(err)
		os.Exit(1)
	}
}

func (what *Base) Init(buildTimestamp string) *Base {
	what.buildTimestamp = buildTimestamp
	return what.initRoot().initAnalyze().initCreate().initExecute().initExplain().initList().initPrint().initQuit().initServer().initVersion()
}
