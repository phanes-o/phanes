package env

import (
	"github.com/spf13/cobra"
)

var (
	boolVarVerbose bool

	Cmd = &cobra.Command{
		Use:   "env",
		Short: "Check or update phanes running environment",
		RunE:  write,
	}
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "phanes env installation",
		RunE:  install,
	}
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Detect phanes env and dependency tools",
		RunE:  check,
	}
)

func init() {
	Cmd.PersistentFlags().BoolVarP(&boolVarVerbose, "verbose", "v", false, "Enable logger output")
	Cmd.AddCommand(installCmd)
	Cmd.AddCommand(checkCmd)
}

func write(cmd *cobra.Command, _ []string) error {
	cmd.Help()
	return nil
}
