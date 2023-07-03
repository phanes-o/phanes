package v1

import (
	"github.com/gin-gonic/gin"
	"hello/bll"
	"hello/model"
	"hello/server/web/middleware"
	"hello/utils"
)

var Manager = &manager{}

func init() {
	RegisterRouter(Manager)
}

type manager struct{}

// Init
func (a *manager) Init(r *gin.RouterGroup) {
	g := r.Group("/manager", middleware.Auth())
	{
		g.POST("/create", a.create)
		g.POST("/update", a.update)
		g.POST("/list", a.list)
		g.POST("/delete", a.delete)
		g.POST("/detail", a.find)
	}
}

// create
func (a *manager) create(c *gin.Context) {
	var (
		in  = &model.ManagerCreateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Manager.Create(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// update
func (a *manager) update(c *gin.Context) {
	var (
		in  = &model.ManagerUpdateRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Manager.Update(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}

// list
func (a *manager) list(c *gin.Context) {
	var (
		in  = &model.ManagerListRequest{}
		out = &model.ManagerListResponse{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Manager.List(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// list
func (a *manager) find(c *gin.Context) {
	var (
		in  = &model.ManagerInfoRequest{}
		out = &model.ManagerInfo{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if out, err = bll.Manager.Find(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, out)
}

// delete
func (a *manager) delete(c *gin.Context) {
	var (
		in  = &model.ManagerDeleteRequest{}
		err error
	)

	if err = c.ShouldBindJSON(in); err != nil {
		c.Error(err)
		return
	}

	if err = bll.Manager.Delete(c.Request.Context(), in); err != nil {
		c.Error(err)
		return
	}
	utils.ResponseOk(c, nil)
}
