package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
)

type GetSomeProducts struct{}

func (gsp GetSomeProducts) Do(c *gin.Context) {
	page, ok := readPagination(c)
	if !ok {
		return
	}

	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	ps, ok := gsp.getFromDB(c, repo, page)
	if !ok {
		return
	}

	ps = gsp.generateDiscount(ps)

	c.JSON(http.StatusOK, ps)
}

func (gsp GetSomeProducts) getFromDB(c *gin.Context, repo database.ProductRepository, page int) (ps []*product.Product, ok bool) {
	ps, err := repo.Get(page, c.GetInt("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (gsp GetSomeProducts) generateDiscount(psR []*product.Product) (ps []*product.Product) {
	var discountGenerator product.DiscountGenerator
	ps = make([]*product.Product, len(psR))
	for i, p := range psR {
		var vault, port int
		if p.Vault != nil {
			vault = *p.Vault
		}
		if p.Port != nil {
			port = *p.Port
		}
		discountGenerator = product.NewDiscountGenerator(vault, port, *p.Quantity, *p.ShippingPrice)
		p.Discount = discountGenerator.Generate()
		ps[i] = p
	}
	return
}
