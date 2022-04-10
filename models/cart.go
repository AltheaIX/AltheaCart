package models

import (
	"AlthCart/config"
	"fmt"
)

type Carts struct {
	Id       int64 `json:"id"`
	Product  Products
	UserId   int64 `json:"userId" db:"user_id"`
	Quantity int64 `json:"quantity"`
}

func GetUserCart(id int64) ([]Carts, error) {
	db := config.OpenConn()
	defer db.Close()

	data := []Carts{}
	stmt, err := db.Preparex("SELECT * FROM users_cart WHERE user_id=$1")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Queryx(id)
	if err != nil {
		panic(err)
	}
	var pId int64

	for rows.Next() {
		row := Carts{}
		rows.Scan(&row.Id, &pId, &row.UserId, &row.Quantity)
		row.Product, err = GetProductById(pId)
		if err != nil {
			panic(err)
		}
		data = append(data, row)
	}

	return data, nil
}

func TestGetUserCart() {
	data, err := GetUserCart(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
}
