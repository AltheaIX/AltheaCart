package models

import (
	"AlthCart/config"
	"database/sql"
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

func AddCarts(id int64, uId int64) error {
	db := config.OpenConn()
	defer db.Close()
	var quantity int

	_, err := GetProductById(id)
	if err != nil {
		return err
	}

	err = db.Get(&quantity, "SELECT quantity FROM users_cart WHERE user_id=$1 AND product_id=$2", uId, id)
	if err != nil {
		if err == sql.ErrNoRows {
			db.Queryx("INSERT INTO users_cart(user_id, product_id, quantity) VALUES ($1, $2, 1);", uId, id)
			return nil
		}
		return err
	}

	stmt, err := db.Preparex("UPDATE users_cart SET quantity=quantity + 1 WHERE user_id=$1 AND product_id=$2")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = stmt.Queryx(id, uId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func RemoveCarts(id int64, uId int64) error {
	db := config.OpenConn()
	defer db.Close()

	var quantity int
	stmt, err := db.Preparex("SELECT quantity FROM users_cart WHERE id=$1 AND user_id=$2")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = stmt.Get(&quantity, id, uId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if quantity > 1 {
		stmt, err := db.Preparex("UPDATE users_cart SET quantity=quantity - 1 WHERE user_id=$1 AND product_id=$2")
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		_, err = stmt.Queryx(uId, id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		return nil
	}

	stmt, err = db.Preparex("DELETE FROM users_cart WHERE user_id=$1 AND product_id=$2")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt.Queryx(uId, id)
	return nil
}

func TestRemoveCarts() {
	RemoveCarts(2, 1)
}

func TestAddCarts() {
	AddCarts(1, 1)
}

func TestGetUserCart() {
	data, err := GetUserCart(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
}
