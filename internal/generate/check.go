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

func checkDestination(src, conf *Destinations) *Destinations {
	if conf.Bll != "" {
		src.Bll = conf.Bll
	}
	if conf.Model != "" {
		src.Model = conf.Model
	}
	if conf.Entity != "" {
		src.Entity = conf.Entity
	}
	if conf.Cache != "" {
		src.Cache = conf.Cache
	}
	if conf.HttpApi != "" {
		src.HttpApi = conf.HttpApi
	}
	if conf.GrpcApi != "" {
		src.GrpcApi = conf.GrpcApi
	}
	if conf.StoreMysql != "" {
		src.StoreMysql = conf.StoreMysql
	}
	if conf.StorePostgres != "" {
		src.StorePostgres = conf.StorePostgres
	}
	if conf.StoreInterface != "" {
		src.StoreInterface = conf.StoreInterface
	}
	return src
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
