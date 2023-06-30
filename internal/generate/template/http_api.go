package template

func init() {
	register(HttpApiTemplate, httpApi)
}

var httpApi = `

package v1

import (
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/bll"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/server/web/middleware"
	"{{.ProjectName}}/utils"
)

var {{.StructName}} = &{{.CamelName}}{}

func init() {
	RegisterRouter({{.StructName}})
}


type {{.CamelName}} struct {}

// Init 
func (a *{{.CamelName}}) Init (r *gin.RouterGroup) {
	g := r.Group("/{{.CamelName}}",  middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
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
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

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

	if err = c.ShouldBindJSON(in); err != nil {
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
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

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
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if  err = bll.{{.StructName}}.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
`
