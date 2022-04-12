package api

import (
	"fmt"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	fmt.Println(products)
<<<<<<< HEAD
	for _, Barang := range products{
		response.Products = append(response.Products, Product{
			Name: Barang.ProductName,
			Price: Barang.Price,
			Category: Barang.Category, 
		})
	}
	encoder.Encode(ProductListSuccessResponse{Products: response.Products})
	// TODO: replace this

}
=======

	encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
}
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
