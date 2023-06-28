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

var {{.TitleName}} = &{{.Name}}{}

func init() {
	RegisterRouter({{.TitleName}})
}


type {{.Name}} struct {}

// Init 
func (a *{{.Name}}) Init (r *gin.RouterGroup) {
	g := r.Group("/{{.Name}}",  middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create 
func (a *{{.Name}}) create(c *gin.Context) {
	var (
		in  = &model.{{.TitleName}}CreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.{{.TitleName}}.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update 
func (a *{{.Name}}) update(c *gin.Context) {
	var (
		in  = &model.{{.TitleName}}UpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.{{.TitleName}}.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list 
func (a *{{.Name}}) list(c *gin.Context) {
	var (
		in  = &model.{{.TitleName}}ListRequest{}
		out  = &model.{{.TitleName}}ListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.{{.TitleName}}.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list 
func (a *{{.Name}}) find(c *gin.Context) {
	var (
		in  = &model.{{.TitleName}}InfoRequest{}
		out  = &model.{{.TitleName}}Info{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.{{.TitleName}}.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete 
func (a *{{.Name}}) delete(c *gin.Context) {
	var (
		in  = &model.{{.TitleName}}DeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if  err = bll.{{.TitleName}}.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
`
