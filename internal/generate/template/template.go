package template

var template = make(map[Type]string)

type Type string

const (
	BllTemplate            Type = "bll"
	ModelTemplate          Type = "model"
	EntityTemplate         Type = "entity"
	HttpApiTemplate        Type = "http_api"
	GrpcApiTemplate        Type = "grpc_api"
	StoreInterfaceTemplate Type = "store_interface"
	StoreMysqlTemplate     Type = "store_postgres"
	StorePostgresTemplate  Type = "store_mysql"
)

func register(key Type, val string) {
	template[key] = val
}

func GetTemplate() map[string]string {
	return template
}
