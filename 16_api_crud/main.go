package main

import (
	"16_api_crud/api/handlers"
	"16_api_crud/infrastructure/repository"
	"16_api_crud/usecase"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

//func addNewProduct(w http.ResponseWriter, r *http.Request) {
//	var product entities.Product
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	json.Unmarshal(reqBody, &product)
//
//	db := openConnection()
//	_, err := db.Exec("insert into Products (model, company, price) values ($1, $2, $3)",
//		product.Model, product.Company, product.Price)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	defer db.Close()
//
//	json.NewEncoder(w).Encode(product)
//}
//
//func updateProduct(w http.ResponseWriter, r *http.Request) {
//	var product entities.Product
//	id := strings.TrimPrefix(r.URL.Path, "/update-product/")
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	json.Unmarshal(reqBody, &product)
//
//	db := openConnection()
//	_, err := db.Exec("update Products set model = $1, price = $2 where id = $3",
//		product.Model, product.Price, id)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	defer db.Close()
//
//	json.NewEncoder(w).Encode(product)
//}

//func getAllProducts(w http.ResponseWriter, r *http.Request) {
//	var products []entities.Product
//	db := openConnection()
//	rows, err := db.Query("select * from Products")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		p := entity.Product{}
//		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		products = append(products, p)
//	}
//	defer db.Close()
//
//	json.NewEncoder(w).Encode(products)
//}

//func getProductById(w http.ResponseWriter, r *http.Request) {
//	var product entities.Product
//	id := strings.TrimPrefix(r.URL.Path, "/get-product/")
//
//	db := openConnection()
//	row := db.QueryRow("select * from Products where id = $1", id)
//	err := row.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	defer db.Close()
//
//	json.NewEncoder(w).Encode(product)
//}

//func openConnection() *sql.DB {
//	connStr := "user=postgres password=postgres dbname=go-test sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	return db
//}

func handleRequests() {
	http.HandleFunc("/", homePage)
	//http.HandleFunc("/add-product", addNewProduct)
	//http.HandleFunc("/update-product/", updateProduct)
	//http.HandleFunc("/get-product/", getProductById)
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Panicln(err)
	}
}

func main() {
	connStr := "user=postgres password=postgres dbname=go-test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	productRepo := repository.NewProductPG(db)
	productService := usecase.NewService(productRepo)

	handlers.MakeAllProductHandler("/get-all-products", productService)
	handleRequests()
}
