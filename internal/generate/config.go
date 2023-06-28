package generate

import (
	"fmt"
)

type PathName string

const (
	BllName            PathName = "bll"
	ModelName                   = "model"
	CacheName                   = "cache"
	EntityName                  = "entity"
	HttpApiName                 = "http_api"
	GrpcApiName                 = "grpc_api"
	StoreMysqlName              = "store_mysql"
	StorePostgresName           = "store_postgres"
	StoreInterfaceName          = "store_interface"
)

type GenType string

const (
	GenTypeBll           = "bll"
	GenTypeModel         = "model"
	GenTypeEntity        = "entity"
	GenTypeHttpApi       = "api.http"
	GenTypeGrpcApi       = "api.grpc"
	GenTypeApiAll        = "api.all"
	GenTypeStoreMysql    = "store.mysql"
	GenTypeStorePostgres = "store.postgres"
	GenTypeStoreAll      = "store.all"
)

type Destinations struct {
	Bll            string `yaml:"bll"`
	Model          string `yaml:"model"`
	Cache          string `yaml:"cache"`
	Entity         string `yaml:"entity"`
	HttpApi        string `yaml:"http_api"`
	GrpcApi        string `yaml:"grpc_api"`
	StoreMysql     string `yaml:"store_mysql"`
	StorePostgres  string `yaml:"store_postgres"`
	StoreInterface string `yaml:"store_interface"`
}

func DefaultDestinations(project string, pwd string) *Destinations {
	var template = "%s/%s/%s"
	return &Destinations{
		Bll:            fmt.Sprintf(template, pwd, project, "bll"),
		Model:          fmt.Sprintf(template, pwd, project, "model"),
		Cache:          fmt.Sprintf(template, pwd, project, "store/redis"),
		Entity:         fmt.Sprintf(template, pwd, project, "model/entity"),
		HttpApi:        fmt.Sprintf(template, pwd, project, "server/grpc/v1"),
		GrpcApi:        fmt.Sprintf(template, pwd, project, "setver/web/v1"),
		StoreMysql:     fmt.Sprintf(template, pwd, project, "store/mysql"),
		StorePostgres:  fmt.Sprintf(template, pwd, project, "store/postgres"),
		StoreInterface: fmt.Sprintf(template, pwd, project, "store"),
	}
}

// todo: path merge
func resolvePaths(path map[PathName]string, destinations *Destinations) map[PathName]string {
	if _, ok := path[BllName]; !ok {
		path[BllName] = destinations.Bll
	}

	if _, ok := path[ModelName]; !ok {
		path[ModelName] = destinations.Model
	}

	if _, ok := path[CacheName]; !ok {
		path[CacheName] = destinations.Cache
	}

	if _, ok := path[CacheName]; !ok {
		path[CacheName] = destinations.Cache
	}

	if _, ok := path[EntityName]; !ok {
		path[EntityName] = destinations.Entity
	}

	if _, ok := path[HttpApiName]; !ok {
		path[HttpApiName] = destinations.HttpApi
	}

	if _, ok := path[GrpcApiName]; !ok {
		path[GrpcApiName] = destinations.GrpcApi
	}

	if _, ok := path[StoreMysqlName]; !ok {
		path[StoreMysqlName] = destinations.HttpApi
	}

	if _, ok := path[StorePostgresName]; !ok {
		path[StorePostgresName] = destinations.StorePostgres
	}

	if _, ok := path[StoreInterfaceName]; !ok {
		path[StoreInterfaceName] = destinations.StoreInterface
	}

	return path
}
