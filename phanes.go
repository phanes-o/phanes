package main

import (
	"log"

	"github.com/phanes-o/phanes/internal/env"
	"github.com/phanes-o/phanes/internal/generate"
	"github.com/phanes-o/phanes/internal/project"
	"github.com/phanes-o/phanes/internal/proto"
	"github.com/phanes-o/phanes/internal/run"
	"github.com/phanes-o/phanes/internal/upgrade"

	"github.com/spf13/cobra"
)

// release is the current phanes tool version.
const release = "v0.1.2"

var rootCmd = &cobra.Command{
	Use:     "phanes",
	Short:   "phanes: An elegant toolkit for Go microservices.",
	Long:    `phanes: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(env.Cmd)
	rootCmd.AddCommand(generate.Cmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
