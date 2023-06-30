package main

import (
	"time"

	"github.com/lib/pq"
)

//dir:bll->./bll/user/user.go
//dir:store_interface->./store/user.go
//dir:http_api->./server/web/v1/user/user.go
// Required: represent this field is required parameter
// AutoGenGormTag: generator will auto generate gorm's tag. if not specify this rule and not specify gorm's or others orm's tag the field's tag will be empty
// NameStyle: you can specify the json tag naming style as snake_case or camelCase or you can directly specify json tag
// EnableValidator: enable http parameter validator, but you must specify validator tag. if you have no this validator tag EnableValidator is invalid
// Except rule tag any others tag will be as fields tag. for example: you use http parameter validator

//project:hello
//generate:bll;store.mysql;store.postgres;api.http;api.grpc;entity;model
type User struct {
	Id        int64          `json:"id" rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator" gorm:"primaryKey;unique;type:BIGINT;autoIncrement;NotNull"`
	Name      string         `json:"name" rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Age       int            `rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Phones    pq.StringArray `rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	CreatedAt time.Time      `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	UpdatedAt time.Time      `rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
}