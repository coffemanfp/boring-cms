package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
)

type GetProduct struct{}

func (gp GetProduct) Do(c *gin.Context) {
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	id, ok := gp.readProductID(c)
	if !ok {
		return
	}

	p, ok := gp.getProductFromDB(c, id, repo)
	if !ok {
		return
	}

	p = gp.generateDiscount(p)

	c.JSON(http.StatusOK, p)
}

func (gp GetProduct) readProductID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (gp GetProduct) getProductFromDB(c *gin.Context, id int, repo database.ProductRepository) (p product.Product, ok bool) {
	p, err := repo.GetOne(id, c.GetInt("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (gp GetProduct) generateDiscount(pr product.Product) (p product.Product) {
	p = pr
	var vault, port int
	if p.Vault != nil {
		vault = *p.Vault
	}
	if p.Port != nil {
		port = *p.Port
	}
	discountGenerator := product.NewDiscountGenerator(vault, port, *p.Quantity, *p.ShippingPrice)
	p.Discount = discountGenerator.Generate()
	return
}
