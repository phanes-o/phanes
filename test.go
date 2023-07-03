package main

import (
	"time"

	"github.com/lib/pq"
)

// Command "generate" specify generate code type. example: api.http
//generate:bll;store.postgres;api.http;entity;model

// Command "project" specify your project to generate
//project:hello

// Command "dir" specify file save path. example: bll->./bll/user/user.go
//dir:bll->./bll/user/user.go
//dir:store_interface->./store/user.go
//dir:http_api->./server/web/v1/user/user.go

// Field Rule tag
// Except rule tag any others tag will be as fields tag. for example: you use http parameter validator

// Parameter: represent this field is a parameter
// Required: represent this field is required parameter
// AutoGenGormTag: generator will auto generate gorm's tag. if not specify this rule and not specify gorm's or others orm's tag the field's tag will be empty
// NameStyle: you can specify the json tag naming style as snake_case or camelCase or you can directly specify json tag
// EnableValidator: enable http parameter validator, but you must specify validator tag. if you have no this validator tag EnableValidator is invalid

//project:hello
type Person struct {
	Id        int64          `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Name      string         `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Arm       string         `rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Age       int            `rule:"AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Phones    pq.StringArray `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	CreatedAt time.Time      `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	UpdatedAt time.Time      `rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
}

//project:hello
type Manager struct {
	Id        int64          `json:"id" rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Name      string         `json:"name" rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Arm       string         `json:"name" rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Age       int            `rule:"AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	Phones    pq.StringArray `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	CreatedAt time.Time      `rule:"Parameter;Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
	UpdatedAt time.Time      `rule:"Parameter;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
}
