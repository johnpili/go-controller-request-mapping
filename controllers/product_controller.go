package controllers

import (
	"github.com/go-zoo/bone"
	"github.com/johnpili/go-controller-request-mapping/models"
	"github.com/psi-incontrol/go-sprocket/sprocket"
	"github.com/shopspring/decimal"
	"net/http"
)

// ProductController ...
type ProductController struct {
	products []models.Product
}

// RequestMapping ...
func (z *ProductController) RequestMapping(router *bone.Mux) {
	router.GetFunc("/v1/get-products", z.GetProducts)
	router.GetFunc("/v1/get-product/:code", z.GetProduct)
}

// NewProductController ...
func NewProductController() *ProductController {
	//region SETUP DUMMY PRODUCTS
	p := make([]models.Product, 0)

	p = append(p, models.Product{
		Code:         "0001",
		Name:         "Product 1",
		PricePerUnit: decimal.NewFromFloat32(99.01),
	})

	p = append(p, models.Product{
		Code:         "0002",
		Name:         "Product 2",
		PricePerUnit: decimal.NewFromFloat32(25.99),
	})

	p = append(p, models.Product{
		Code:         "0003",
		Name:         "Product 3",
		PricePerUnit: decimal.NewFromFloat32(1.25),
	})

	p = append(p, models.Product{
		Code:         "0004",
		Name:         "Product 4",
		PricePerUnit: decimal.NewFromFloat32(2.50),
	})
	//endregion

	return &ProductController{
		products: p,
	}
}

// GetProducts ...
func (z *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	sprocket.RespondOkayJSON(w, z.products)
}

// GetProduct ...
func (z *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	code := bone.GetValue(r, "code")
	if len(code) == 0 {
		sprocket.RespondBadRequestJSON(w, nil)
		return
	}

	for _, item := range z.products {
		if item.Code == code {
			sprocket.RespondOkayJSON(w, item)
			return
		}
	}

	sprocket.RespondNotFoundJSON(w, nil)
}
