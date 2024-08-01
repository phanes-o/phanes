package template

func init() {
	register(HttpApiTemplate, httpApi)
}

var httpApi = `

package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"{{.Module}}/bll"
	"{{.Module}}/model"
	"{{.Module}}/server/web/middleware"
	"{{.Module}}/utils"
)

var {{.StructName}} = &{{.CamelName}}{}

func init() {
	RegisterRouter({{.StructName}})
}


type {{.CamelName}} struct {}

// Init 
func (a *{{.CamelName}}) Init (r *gin.RouterGroup) {
	g := r.Group("/{{.CamelName}}s",  middleware.Auth())
	{
		g.POST("", a.create)
		g.PUT(":id", a.update)
		g.GET("", a.list)
		g.DELETE(":id", a.delete)
		g.GET(":id", a.find)
	}
}

// create 
func (a *{{.CamelName}}) create(c *gin.Context) {
	var (
		in  = &model.{{.StructName}}CreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.{{.StructName}}.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update 
func (a *{{.CamelName}}) update(c *gin.Context) {
	var (
		in  = &model.{{.StructName}}UpdateRequest{}
		err error
		id int64
	)

	// Verify if the ID is a valid positive integer
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0  {
		c.Error(errors.ErrInvalidParams.Error())
		return
	}

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	in.Id = id

	if err = bll.{{.StructName}}.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list 
func (a *{{.CamelName}}) list(c *gin.Context) {
	var (
		in  = &model.{{.StructName}}ListRequest{}
		out  = &model.{{.StructName}}ListResponse{}
		err error
	)

	if err = c.ShouldBindQuery(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.{{.StructName}}.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list 
func (a *{{.CamelName}}) find(c *gin.Context) {
	var (
		in  = &model.{{.StructName}}InfoRequest{}
		out  = &model.{{.StructName}}Info{}
		err error
		id int64
	)

	// Verify if the ID is a valid positive integer
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0  {
		c.Error(errors.ErrInvalidParams.Error())
		return
	}

	in.Id = id

	if out, err = bll.{{.StructName}}.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete 
func (a *{{.CamelName}}) delete(c *gin.Context) {
	var (
		in  = &model.{{.StructName}}DeleteRequest{}
		err error
	    id int64
	)

	// Verify if the ID is a valid positive integer
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0  {
		c.Error(errors.ErrInvalidParams.Error())
		return
	}

	in.Id = id
	if  err = bll.{{.StructName}}.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
`
