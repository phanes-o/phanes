package generate

import (
	"fmt"
	"os"
)

func checkEnvironment() bool {
	// todo: checkEnvironment
	return true
}

func checkConfigFile(configName string) bool {
	var (
		err error
		pwd string
	)
	if pwd, err = os.Getwd(); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s \033[m\n", err)
		return false
	}
	return fileExists(fmt.Sprintf("%s/%s", pwd, configName))
}

func checkProjectExist(project string) bool {
	return fileExists(project)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
