package controller

import (
	"github.com/kataras/iris"
	"gitlab.com/z547743799/iriscontent/service"
)

type PageController struct {
	Ctx     iris.Context
	Service service.ContentService
}

func (c *PageController) Get() {
	list := c.Service.GetContentListByCid(89)
	c.Ctx.ViewData("ad1List", list)
	c.Ctx.View("index.html")
}

func (c *PageController) GetBy(page string) {

	c.Ctx.View(page + ".html")
}
