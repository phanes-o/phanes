package generate

import (
	"fmt"
	"os"
	"strings"
)

func parseCommentGenType(s string) []GenType {
	//codeBuild: bll;store.mysql;store.postgres;api.http;api.grpc;entity;model
	var (
		apiAll   bool
		storeAll bool
		genTypes = make([]GenType, 0)
	)
	split := strings.Split(s, ";")
	if len(split) > 0 {
		for _, t := range split {
			switch t {
			case GenTypeBll:
				genTypes = append(genTypes, GenTypeBll)
			case GenTypeModel:
				genTypes = append(genTypes, GenTypeModel)
			case GenTypeEntity:
				genTypes = append(genTypes, GenTypeEntity)
			case GenTypeGrpcApi:
				genTypes = append(genTypes, GenTypeGrpcApi)
			case GenTypeHttpApi:
				genTypes = append(genTypes, GenTypeHttpApi)
			case GenTypeApiAll:
				apiAll = true
				genTypes = append(genTypes, GenTypeApiAll)
			case GenTypeStoreAll:
				storeAll = true
				genTypes = append(genTypes, GenTypeStoreAll)
			case GenTypeStoreMysql:
				genTypes = append(genTypes, GenTypeStoreMysql)
			case GenTypeStorePostgres:
				genTypes = append(genTypes, GenTypeStorePostgres)
			}
		}
	}
	for i, v := range genTypes {
		if apiAll && (v == GenTypeHttpApi || v == GenTypeGrpcApi) {
			tmp := append(genTypes[:i], genTypes[i+1:]...)
			genTypes = tmp
		}
		if storeAll && (v == GenTypeStoreMysql || v == GenTypeStorePostgres) {
			tmp := append(genTypes[:i], genTypes[i+1:]...)
			genTypes = tmp
		}
	}
	return genTypes
}

func parseCommentDir(val string) (PathName, string) {
	split := strings.Split(val, "->")
	if len(split) != 2 {
		fmt.Fprint(os.Stderr, "\033[31mERROR: your dir error  \033[m\n")
		os.Exit(0)
	}

	return PathName(split[0]), split[1]
}
