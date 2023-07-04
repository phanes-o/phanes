package generate

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
)

type PathName string

const (
	BllName            PathName = "bll"
	ModelName                   = "model"
	CacheName                   = "cache"
	EntityName                  = "entity"
	MappingName                 = "mapping"
	HttpApiName                 = "http_api"
	GrpcApiName                 = "grpc_api"
	StoreMysqlName              = "store_mysql"
	StorePostgresName           = "store_postgres"
	StoreInterfaceName          = "store_interface"
)

type GenType string

const (
	GenTypeBll           GenType = "bll"
	GenTypeModel                 = "model"
	GenTypeEntity                = "entity"
	GenTypeHttpApi               = "api.http"
	GenTypeGrpcApi               = "api.grpc"
	GenTypeApiAll                = "api.all"
	GenTypeStoreMysql            = "store.mysql"
	GenTypeStorePostgres         = "store.postgres"
)

func destinations(project string, pwd string) map[PathName]string {
	var template = "%s/%s/%s"
	return map[PathName]string{
		BllName:            fmt.Sprintf(template, pwd, project, "bll"),
		ModelName:          fmt.Sprintf(template, pwd, project, "model"),
		CacheName:          fmt.Sprintf(template, pwd, project, "store/redis"),
		EntityName:         fmt.Sprintf(template, pwd, project, "model/entity"),
		MappingName:        fmt.Sprintf(template, pwd, project, "model/mapping"),
		HttpApiName:        fmt.Sprintf(template, pwd, project, "server/web/v1"),
		GrpcApiName:        fmt.Sprintf(template, pwd, project, "server/grpc/v1"),
		StoreMysqlName:     fmt.Sprintf(template, pwd, project, "store/mysql"),
		StorePostgresName:  fmt.Sprintf(template, pwd, project, "store/postgres"),
		StoreInterfaceName: fmt.Sprintf(template, pwd, project, "store"),
	}
}

// /server/web/v1/user/user.go
func resolvePaths(project, pwd string, structName StructName, paths map[PathName]string, destinations map[PathName]string) map[PathName]string {
	var (
		ok       bool
		fileName string
		suffix   = fmt.Sprintf("%s%s", Camel2Case(string(structName)), ".go")
	)

	for name, addr := range destinations {
		if fileName, ok = paths[name]; !ok {
			// use default destination
			fileName = addr
			fileName = path.Join(fileName, suffix)
		} else {
			// user specify destination
			if !strings.Contains(fileName, project) {
				fmt.Println(color.RedString("Error: Your path is not your project"))
				os.Exit(1)
			}
			if strings.Contains(fileName, ".") {
				fileName = strings.TrimLeft(fileName, ".")

				if !strings.Contains(fileName, suffix) {
					fileName = path.Join(fileName, suffix)
				}
				fileName = path.Join(pwd, fileName)
			} else {
				if !strings.Contains(fileName, suffix) {
					fileName = path.Join(fileName, suffix)
				}
			}
		}

		paths[name] = fileName
	}
	return paths
}
