package handlers

import (
	"16_api_crud/api/middleware"
	"encoding/json"
	"net/http"

	"16_api_crud/api/presenter"
	"16_api_crud/entities"
	"16_api_crud/usecase"
)

func getAllProducts(service usecase.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading products"
		var products []*entities.Product
		var err error

		products, err = service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		var toJson []*presenter.Product
		for _, p := range products {
			toJson = append(toJson, &presenter.Product{
				Id:      p.Id,
				Model:   p.Model,
				Company: p.Company,
				Price:   p.Price,
			})
		}

		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	}
}

////MakeProductHandlers make url handlers
//func MakeProductHandlers(r *mux.Router, n negroni.Negroni, service usecase.Getter) {
//	r.Handle("/get-all-products", n.With(
//		negroni.Wrap(getAllProducts(service)),
//	)).Methods("GET").Name("getAllProducts")
//}

//MakeAllProductHandler sets handler behavior for url pattern
func MakeAllProductHandler(pattern string, service usecase.Getter) {
	http.HandleFunc(pattern, middleware.Cors(getAllProducts(service)))
}

//func MakeAllProductHandler(pattern string, service usecase.Getter) {
//	http.HandleFunc(pattern, getAllProducts(service))
//}
//func MakeAllProductHandler(pattern string, service usecase.Getter) {
//	http.HandleFunc(pattern, getAllProducts(service))
//}
//func MakeAllProductHandler(pattern string, service usecase.Getter) {
//	http.HandleFunc(pattern, getAllProducts(service))
//}
