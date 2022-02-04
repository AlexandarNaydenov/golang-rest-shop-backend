package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	. "github.com/golang-rest-shop-backend/pkg/structs"
	"github.com/nu7hatch/gouuid"
	"os"
)

var db *sql.DB

func InitMySqlConnection() error {

	config := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_IP_ADDRESS"),
		DBName: "products",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return fmt.Errorf("database opening failed with error: %s", err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return fmt.Errorf("ping my sql failed with error: %s", err.Error())
	}

	return nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error while reading all products from database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Category, &p.Quantity, &p.Price); err != nil {
			return nil, fmt.Errorf("parsing to a product failed with: %v", err)
		}
		products = append(products, p)
	}

	return products, nil
}

func GetProductById(productId string) (*Product, error) {
	row := db.QueryRow("SELECT * FROM products WHERE id = ?", productId)

	var p Product
	if err := row.Scan(&p.ID, &p.Title, &p.Category, &p.Quantity, &p.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no product with id: %s", productId)
		}
		return nil, fmt.Errorf("searching for %s failed with: %s", productId, err)
	}

	return &p, nil
}

func GetAllOrders() ([]Order, error) {
	var orders []Order

	rows, _ := db.Query("SELECT * FROM orders")
	defer rows.Close()

	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.Name, &o.Address, &o.Phone, &o.Products, &o.Price, &o.Status); err != nil {
			return nil, fmt.Errorf("getting all products failed with: %v", err)
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func GetOrderById(orderId string) (*Order, error) {
	row := db.QueryRow("SELECT * FROM orders WHERE id = ?", orderId)

	var o Order
	if err := row.Scan(&o.ID, &o.Name, &o.Address, &o.Phone, &o.Products, &o.Price, &o.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order with id: %s", orderId)
		}
		return nil, fmt.Errorf("searching for %s failed with: %s", orderId, err)
	}

	return &o, nil
}

func AddOrder(order *Order) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("failed to generate uuid error: %s", err)
	}

	products := ""
	for _, p := range order.Products {
		products += p.ID + ","
	}

	status := "Accepted"

	_, err = db.Query("INSERT INTO orders (ID, ID, Address, Phone, Products, Price, Status) VALUES (?,?,?,?,?,?)", id.String(), order.Name, order.Address, order.Phone, products, status)
	if err != nil {
		return "", fmt.Errorf("failed to add order to the database, error: %s", err)
	}

	return id.String(), nil
}

func DeleteOrder(orderId int) error {
	result, err := db.Exec("DELETE FROM orders WHERE ID = ?;", orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order from the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no order with id: %d", orderId)
	}

	return nil
}

func ChangeProductQuantity(productId string, quantity int) error {
	var p Product

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", productId)
	if err := row.Scan(&p.ID, &p.Title, &p.Category, &p.Quantity, &p.Price); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no product with id: %s", productId)
		}
		return fmt.Errorf("searching for id: %s failed with: %s", productId, err)
	}

	newQuantity := p.Quantity - quantity
	if newQuantity < 0 {
		return fmt.Errorf("not enough quantity of product: %s", p.Title)
	}

	if _, err := db.Query("UPDATE products SET quantity = ? WHERE id = ?", newQuantity, p.ID); err != nil {
		return fmt.Errorf("updating quantity failed with: %s", err)
	}

	return nil
}
