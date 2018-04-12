package controllers

import (
	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type App struct {
	gormc.Controller
}

var (
	url string = "bento/dist/index.html"
)

func (c App) Index() revel.Result {
	var products []*models.Product
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&products)
	c.DB.Raw("SELECT * FROM product;").Scan(&products)
	//return c.Render(products)
	return c.RenderFileName(url, revel.NoDisposition)
}

func (c App) Help() revel.Result {
	return c.Render()
}
