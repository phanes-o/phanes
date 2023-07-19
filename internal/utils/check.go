package utils

import (
	"os"
)

func CheckEnvironment() bool {
	// todo: checkEnvironment
	return true
}
func CheckProjectExist(project string) bool {
	return FileExists(project)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
