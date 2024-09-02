package base

import (
	"fmt"
	"github.com/spf13/cobra"
)

const ()

func (what *Base) initRoot() *Base {
	what.rootCmd = &cobra.Command{
		Use:           "base",
		Version:       BaseVersion,
		Short:         "\n" + Logo,
		Long:          "\n" + Logo + "\n\n" + fmt.Sprintf(VersionText, what.buildTimestamp) + "\n\n",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run:           what.run,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	return what
}

func (what *Base) run(cmd *cobra.Command, _ []string) {
	if !what.flags.interactiveFlag {
		cmd.Println("Please add the --interactive flag to run in interactive mode.")
		return
	}
}

func (what *Base) readConfig(cmd *cobra.Command, buildTimestamp string) *Config {
	cfg := new(Config).Defaults(buildTimestamp)
	// MORE
	return cfg
}
