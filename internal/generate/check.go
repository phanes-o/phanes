package generate

import (
	"os"
)

func checkEnvironment() bool {
	// todo: checkEnvironment
	return true
}
func checkProjectExist(project string) bool {
	return fileExists(project)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
