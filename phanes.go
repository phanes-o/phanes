package main

import (
	"log"

	"github.com/lizhiqpxv/phanes/internal/env"
	"github.com/lizhiqpxv/phanes/internal/project"
	"github.com/lizhiqpxv/phanes/internal/proto"
	"github.com/lizhiqpxv/phanes/internal/run"
	"github.com/lizhiqpxv/phanes/internal/upgrade"

	"github.com/spf13/cobra"
)

// release is the current phanes tool version.
const release = "v0.1.1"

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

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
