package generate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Cmd represents the new command.
var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "codeBuild all business logic code",
	Long:  "codeBuild business logic cod include bll, store, http api, grpc api, entity",
	Run:   run,
}

var (
	configName     string
	workspace      bool
	importExistMap = map[string]struct{}{}
)

func init() {
	Cmd.Flags().BoolVarP(&workspace, "workspace", "w", false, "enable workspace")
	Cmd.Flags().StringVarP(&configName, "config", "c", "generator.go", "codeBuild config file")
}

func run(cmd *cobra.Command, args []string) {
	var (
		err       error
		generator *Generator
	)
	// Check args

	// If enable workspace. check workspace environment. means Phanes command must run in go workspace
	if workspace {
		if !checkEnvironment() {
			fmt.Fprint(os.Stderr, "\033[31mERROR: can not detection go workspace. please run `go work init` to init go workspace\033[m\n")
			return
		}
	}

	// Build Generator
	// Parse template and replace it
	if generator, err = ReadSource(configName); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Generate failed %s \033[m\n", err)
		return
	}
	// Generate nad write code to destination file.
	if err = generator.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Generate failed %s \033[m\n", err)
		return
	}
}
