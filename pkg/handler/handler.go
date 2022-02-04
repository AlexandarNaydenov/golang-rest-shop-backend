package handler

import (
	"encoding/json"
	"fmt"
	"github.com/golang-rest-shop-backend/pkg/service"
	"github.com/golang-rest-shop-backend/pkg/structs"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func GetAllProductHandler(w http.ResponseWriter, r *http.Request) {
	currency := r.FormValue("currency")

	bytes, err := service.GetAllProducts(currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["productId"]
	currency := r.FormValue("currency")

	bytes, err := service.GetProductById(productId, currency)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func GetAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	currency := r.FormValue("currency")

	bytes, err := service.GetAllOrders(currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["orderId"]
	currency := r.FormValue("currency")

	bytes, err := service.GetOrderById(orderId, currency)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func AddOrderHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var order structs.Order
	err := decoder.Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "request body has wrong format: %s\n", err)
		return
	}

	orderID, err := service.AddOrder(&order)
	if err != nil {
		if strings.HasPrefix(err.Error(), "not enough quantity") {
			fmt.Fprint(w, err)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "Successful purchase: %s", orderID)
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
}

func ChangeOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["orderId"]
	id, err := strconv.Atoi(orderId)
	if err != nil {
		fmt.Fprintf(w, "request id has wrong format: %s\n", err)
		return
	}

	err = service.DeleteOrder(id)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "order %d deleted", id)
}
