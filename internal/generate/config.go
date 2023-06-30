package generate

import (
	"fmt"
	"strings"
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

type Destinations struct {
	Bll            string
	Model          string
	Cache          string
	Entity         string
	Mapping        string
	HttpApi        string
	GrpcApi        string
	StoreMysql     string
	StorePostgres  string
	StoreInterface string
}

func destinations(project string, pwd string) map[PathName]string {
	var template = "%s/%s/%s"
	return map[PathName]string{
		BllName:            fmt.Sprintf(template, pwd, project, "bll"),
		ModelName:          fmt.Sprintf(template, pwd, project, "model"),
		CacheName:          fmt.Sprintf(template, pwd, project, "store/redis"),
		EntityName:         fmt.Sprintf(template, pwd, project, "model/entity"),
		MappingName:        fmt.Sprintf(template, pwd, project, "model/mapping"),
		HttpApiName:        fmt.Sprintf(template, pwd, project, "server/grpc/v1"),
		GrpcApiName:        fmt.Sprintf(template, pwd, project, "server/web/v1"),
		StoreMysqlName:     fmt.Sprintf(template, pwd, project, "store/mysql"),
		StorePostgresName:  fmt.Sprintf(template, pwd, project, "store/postgres"),
		StoreInterfaceName: fmt.Sprintf(template, pwd, project, "store"),
	}
}

func DefaultDestinations(project string, pwd string) *Destinations {
	var template = "%s/%s/%s"
	return &Destinations{
		Bll:            fmt.Sprintf(template, pwd, project, "bll"),
		Model:          fmt.Sprintf(template, pwd, project, "model"),
		Cache:          fmt.Sprintf(template, pwd, project, "store/redis"),
		Entity:         fmt.Sprintf(template, pwd, project, "model/entity"),
		Mapping:        fmt.Sprintf(template, pwd, project, "model/mapping"),
		HttpApi:        fmt.Sprintf(template, pwd, project, "server/grpc/v1"),
		GrpcApi:        fmt.Sprintf(template, pwd, project, "setver/web/v1"),
		StoreMysql:     fmt.Sprintf(template, pwd, project, "store/mysql"),
		StorePostgres:  fmt.Sprintf(template, pwd, project, "store/postgres"),
		StoreInterface: fmt.Sprintf(template, pwd, project, "store"),
	}
}

// /server/web/v1/user/user.go
func resolvePaths(project, pwd string, structName StructName, paths map[PathName]string, destinations map[PathName]string) map[PathName]string {
	var (
		ok       bool
		fileName string
	)

	for name, addr := range destinations {
		if fileName, ok = paths[name]; !ok {
			fileName = addr
		}
		prefix := fmt.Sprintf("%s/%s", pwd, project)
		if !strings.Contains(fileName, prefix) {
			fileName = prefix + strings.TrimLeft(fileName, ".")
		}

		fileName = strings.TrimRight(fileName, "/")

		suffix := fmt.Sprintf("%s%s", Camel2Case(string(structName)), ".go")
		if !strings.Contains(fileName, suffix) {
			fileName += "/" + suffix
		}
		paths[name] = fileName
	}
	return paths
}
