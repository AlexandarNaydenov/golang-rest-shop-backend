package main

import (
	"github.com/golang-rest-shop-backend/pkg/database"
	"github.com/golang-rest-shop-backend/pkg/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	err := database.InitMySqlConnection()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", handler.GetAllProductHandler).Methods("GET")
	r.HandleFunc("/orders", handler.GetAllOrdersHandler).Methods("GET")
	r.HandleFunc("/product/{productId}", handler.GetProductHandler).Methods("GET")
	r.HandleFunc("/order/{orderId}", handler.GetOrderHandler).Methods("GET")

	r.HandleFunc("/order", handler.AddOrderHandler).Methods("POST")
	r.HandleFunc("/product", handler.AddProductHandler).Methods("POST")

	r.HandleFunc("/order/changeStatus", handler.ChangeOrderStatusHandler).Methods("PUT")

	r.HandleFunc("/delete/product/{orderId}", handler.DeleteProductHandler).Methods("DELETE")
	r.HandleFunc("/delete/order/{orderId}", handler.DeleteOrderHandler).Methods("DELETE")

	log.Println("Listening to port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
