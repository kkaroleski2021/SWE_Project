package service

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/golang-rest-shop-backend/pkg/database"
	. "github.com/golang-rest-shop-backend/pkg/structs"
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

func GetAllProducts(currency string) ([]Product, error) {
	products, err := database.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products with error: %s\n", err)
	}

	for i := range products {
		err := convertPrice(&products[i], currency)
		if err != nil {
			return nil, err
		}
	}

	return products, nil
}

func GetProductById(id string, currency string) (*Product, error) {
	product, err := database.GetProductById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find such product error: %s\n", err)
	}

	if err = convertPrice(&product, currency); err != nil {
		return nil, err
	}

	return product, nil
}

func GetAllOrders(currency string) ([]Order, error) {
	orders, err := database.GetAllOrders()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products with error: %s\n", err)
	}

	for i := range orders {
		if err = convertPrice(&orders[i], currency); err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func GetOrderById(id string, currency string) (*Order, error) {
	order, err := database.GetOrderById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find such order error: %s\n", err)
	}

	if err = convertPrice(&order, currency); err != nil {
		return nil, err
	}

	return order, nil
}

func AddOrder(order *Order) (string, error) {
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
		err = database.AddOrderedProduct(&OrderedProduct{
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

func AddProduct(product *Product) (string, error) {
	productId, err := database.AddProduct(product)
	if err != nil {
		return "", err
	}

	return productId, nil
}

func UpdateProduct(product *Product) error {
	err := database.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrder(order *Order) error {
	err := database.UpdateOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(orderId string) error {
	if err := database.DeleteAllProductsForAnOrder(orderId); err != nil {
		return err
	}

	if err := database.DeleteOrder(orderId); err != nil {
		return err
	}

	return nil
}

func DeleteProduct(productId string) error {
	if err := database.DeleteProduct(productId); err != nil {
		return err
	}

	return nil
}

func convertPrice(object interface{}, currency string) error {
	if currency == "" {
		return nil
	}

	rate, err := getRates(currency)
	if err != nil {
		return err
	}

	switch v := object.(type) {
	case *Order:
		{
			v.Price = math.Round(rate*v.Price*100) / 100
			for i := range v.Products {
				v.Products[i].Price = math.Round(rate*v.Products[i].Price*100) / 100
			}
		}
	case *Product:
		{
			v.Price = math.Round(rate*v.Price*100) / 100
		}
	default:
		return fmt.Errorf("unsupported type")
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
