package v1

import (
	"github.com/gin-gonic/gin"
	"hello/bll"
	"hello/model"
	"hello/server/web/middleware"
	"hello/utils"
)

var Person = &person{}

func init() {
	RegisterRouter(Person)
}

type person struct{}

// Init
func (a *person) Init(r *gin.RouterGroup) {
	g := r.Group("/person", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *person) create(c *gin.Context) {
	var (
		in  = &model.PersonCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Person.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *person) update(c *gin.Context) {
	var (
		in  = &model.PersonUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Person.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *person) list(c *gin.Context) {
	var (
		in  = &model.PersonListRequest{}
		out = &model.PersonListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Person.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *person) find(c *gin.Context) {
	var (
		in  = &model.PersonInfoRequest{}
		out = &model.PersonInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Person.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *person) delete(c *gin.Context) {
	var (
		in  = &model.PersonDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Person.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
