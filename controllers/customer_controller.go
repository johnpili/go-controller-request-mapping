package controllers

import (
	"github.com/go-zoo/bone"
	"net/http"
)

// CustomerController ...
type CustomerController struct{}

// RequestMapping ...
func (z *CustomerController) RequestMapping(router *bone.Mux) {
	router.GetFunc("/v1/get-customers", z.GetCustomers)
	router.GetFunc("/v1/get-customer/:code", z.GetCustomer)
	router.GetFunc("/v1/get-customer-orders/:code", z.GetCustomerOrders)
}

// GetCustomers ...
func (z *CustomerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
}

// GetCustomer ...
func (z *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
}

// GetCustomerOrders
func (z *CustomerController) GetCustomerOrders(w http.ResponseWriter, r *http.Request) {
}