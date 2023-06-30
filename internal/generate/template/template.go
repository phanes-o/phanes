package template

var template = make(map[Type]string)

type Type string

const (
	BllTemplate            Type = "bll"
	ModelTemplate          Type = "model"
	MappingTemplate        Type = "mapping"
	EntityTemplate         Type = "entity"
	HttpApiTemplate        Type = "http_api"
	GrpcApiTemplate        Type = "grpc_api"
	StoreInterfaceTemplate Type = "store_interface"
	StoreMysqlTemplate     Type = "store_mysql"
	StorePostgresTemplate  Type = "store_postgres"
)

func register(key Type, val string) {
	template[key] = val
}

func Template() map[Type]string {
	return template
}

func Get(t Type) string {
	if tmpl, ok := template[t]; ok {
		return tmpl
	}
	return ""
}
