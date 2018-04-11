package controllers

import (
	"fmt"

	"github.com/pdmp/amai/app/routes"

	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Products struct {
	gormc.Controller
}

func (c Products) Show() revel.Result {
	var products []*models.Product
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&products)
	c.DB.Raw("SELECT * FROM product;").Scan(&products)
	return c.RenderJSON(products)
}

func (c Products) Crud() revel.Result {
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func (c Products) GetProduct(code string) revel.Result {
	var product models.Product
	c.DB.Raw("SELECT code,price FROM product WHERE code = ?;", code).Scan(&product)
	//return c.RenderJSON(product) //debug
	return c.RenderJSON(product)
}

func (c Products) Delete() revel.Result {
	var products []models.Product
	c.DB.Raw("SELECT product_id,code,price FROM product").Scan(&products)
	//return c.RenderJSON(products) //debugging
	return c.Render(products)
}

func (c Products) Pop(id uint) revel.Result {
	var product models.Product
	c.DB.Raw("DELETE FROM product WHERE product_id = ?;", id).Scan(&product)
	fmt.Println("id", id)
	return c.Redirect("/crud")
	//return c.RenderJSON(product) //debuging
}

func (c Products) Insert(code string, price uint) revel.Result {
	pro := models.Product{Code: code, Price: price}
	c.DB.Create(&pro)
	return c.Redirect(routes.App.Index())
}
