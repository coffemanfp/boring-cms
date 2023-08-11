package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/search"
	"github.com/gin-gonic/gin"
)

type Search struct{}

func (s Search) Do(c *gin.Context) {
	srch, ok := s.readSearch(c)
	if !ok {
		return
	}

	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	ps, ok := s.searchOnDB(c, repo, srch)
	if !ok {
		return
	}

	ps = s.generateDiscount(ps)

	c.JSON(http.StatusOK, ps)
}

func (s Search) readSearch(c *gin.Context) (srch search.Search, ok bool) {
	srch.GuideNumber = c.Query("guideNumber")
	srch.VehiclePlate = c.Query("vehiclePlate")
	srch.Type = c.Query("type")
	startJoinedAt := c.Query("startJoinedAt")
	endJoinedAt := c.Query("endJoinedAt")
	startDeliveredAt := c.Query("startDeliveredAt")
	endDeliveredAt := c.Query("endDeliveredAt")
	srch.Port, ok = readIntFromURL(c, "port", true)
	if !ok {
		return
	}
	srch.Vault, ok = readIntFromURL(c, "vault", true)
	if !ok {
		return
	}
	srch.PriceRange.Start, ok = readFloatFromURL(c, "startPrice", true)
	if !ok {
		return
	}
	srch.PriceRange.End, ok = readFloatFromURL(c, "endPrice", true)
	if !ok {
		return
	}

	var parseTimeValue = func(s string, paramName string) (r time.Time, ok bool) {
		if s == "" {
			ok = true
			return
		}
		r, err := time.Parse(time.RFC3339, s)
		if err != nil {
			err = fmt.Errorf("invalid %s: invalid %s time format of %s", paramName, paramName, s)
			handleError(c, err)
			return
		}
		ok = true
		return
	}

	srch.JoinedAtRange.Start, ok = parseTimeValue(startJoinedAt, "start joined at")
	if !ok {
		return
	}
	srch.JoinedAtRange.End, ok = parseTimeValue(endJoinedAt, "end joined at")
	if !ok {
		return
	}
	srch.DeliveredAtRange.Start, ok = parseTimeValue(startDeliveredAt, "start joined at")
	if !ok {
		return
	}
	srch.DeliveredAtRange.End, ok = parseTimeValue(endDeliveredAt, "end delivered at")
	return
}

func (s Search) searchOnDB(c *gin.Context, repo database.ProductRepository, srch search.Search) (ps []*product.Product, ok bool) {
	ps, err := repo.Search(srch)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (s Search) generateDiscount(psR []*product.Product) (ps []*product.Product) {
	var discountGenerator product.DiscountGenerator
	ps = make([]*product.Product, len(psR))
	for i, p := range psR {
		discountGenerator = product.NewDiscountGenerator(p.Vault, p.Port, p.Quantity, p.ShippingPrice)
		p.Discount = discountGenerator.Generate()
		ps[i] = p
	}
	return
}
