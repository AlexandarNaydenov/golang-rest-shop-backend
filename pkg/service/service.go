package service

import (
	"encoding/json"
	"fmt"
	"github.com/golang-rest-shop-backend/pkg/database"
	"github.com/golang-rest-shop-backend/pkg/structs"
	"math"
	"net/http"
)

type ExchangeRateAPIResponse struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     struct {
		BGN float64 `json:"BGN"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		EUR float64 `json:"EUR"`
		GBP float64 `json:"GBP"`
		USD float64 `json:"USD"`
	} `json:"rates"`
}

func GetAllProducts(currency string) ([]byte, error) {
	p, err := database.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products with error: %s\n", err)
	}

	if currency != "" {
		rate, err := getRates(currency)
		if err != nil {
			return nil, err
		}

		for i := range p {
			p[i].Price = math.Round(p[i].Price*rate*100) / 100
		}
	}

	bytes, _ := json.Marshal(p)
	return bytes, nil
}

func GetProductById(id string, currency string) ([]byte, error) {
	p, err := database.GetProductById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find such product error: %s\n", err)
	}

	if currency != "" {
		rate, err := getRates(currency)
		if err != nil {
			return nil, err
		}

		p.Price = math.Round(rate*p.Price*100) / 100
	}

	bytes, _ := json.Marshal(p)
	return bytes, nil
}

func GetAllOrders(currency string) ([]byte, error) {
	o, err := database.GetAllOrders()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products with error: %s\n", err)
	}

	if currency != "" {
		rate, err := getRates(currency)
		if err != nil {
			return nil, err
		}

		for i := range o {
			o[i].Price = math.Round(o[i].Price*rate*100) / 100
		}
	}

	bytes, _ := json.Marshal(o)
	return bytes, nil
}

func GetOrderById(id string, currency string) ([]byte, error) {
	o, err := database.GetOrderById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find such order error: %s\n", err)
	}

	if currency != "" {
		rate, err := getRates(currency)
		if err != nil {
			return nil, err
		}

		o.Price = math.Round(rate*o.Price*100) / 100
	}

	bytes, _ := json.Marshal(o)
	return bytes, nil
}

func AddOrder(order *structs.Order) (string, error) {
	totalPrice := 0.0

	for _, p := range order.Products {
		err := database.ChangeProductQuantity(p.ID, p.Quantity)
		if err != nil {
			return "", err
		}

		product, err := database.GetProductById(p.ID)
		if err != nil {
			return "", err
		}

		totalPrice += product.Price * float64(p.Quantity)
	}

	order.Price = totalPrice
	order.Status = "Accepted"

	orderId, err := database.AddOrder(order)
	if err != nil {
		return "", err
	}

	for _, p := range order.Products {
		err = database.AddOrderedProduct(&structs.OrderedProduct{
			ProductId:       p.ID,
			ProductQuantity: p.Quantity,
			OrderId:         orderId,
		})
		if err != nil {
			return "", err
		}
	}

	return orderId, nil
}

func DeleteOrder(orderId int) error {
	err := database.DeleteOrder(orderId)
	if err != nil {
		return err
	}

	return nil
}

func getRates(currency string) (float64, error) {
	const accessKey = "a3d5d57407a65c0b4fa4853c2e5cbe07"
	url := fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s&format=1", accessKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("request to exchange rates API failed with error: %s", err)
	}

	decode := json.NewDecoder(resp.Body)
	var exchangeRateResponse ExchangeRateAPIResponse
	err = decode.Decode(&exchangeRateResponse)
	if err != nil {
		return 0, fmt.Errorf("wrong format from exchange rates API, error: %s", err)
	}

	rate := 1.0
	switch currency {
	case "USD":
		rate = exchangeRateResponse.Rates.USD
	case "BGN":
		rate = exchangeRateResponse.Rates.BGN
	case "EUR":
		rate = 1.0
	case "GBP":
		rate = exchangeRateResponse.Rates.GBP
	case "CAD":
		rate = exchangeRateResponse.Rates.CAD
	case "CHF":
		rate = exchangeRateResponse.Rates.CHF
	default:
		return 0, fmt.Errorf("unsupported currency")
	}

	return rate, nil
}
