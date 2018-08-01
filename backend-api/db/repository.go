package db

import (
	"fmt"
	"log"
	"vue-golang-payment-app/backend-api/domain"
)

// SelectAllItems - select all posts
func SelectAllItems() (items domain.Items, err error) {
	stmt, err := Conn.Query("SELECT * FROM items")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer stmt.Close()
	for stmt.Next() {
		var id int64
		var name string
		var description string
		var amount int64

		if err := stmt.Scan(&id, &name, &description, &amount); err != nil {
			fmt.Println(err.Error())
			continue
		}

		item := domain.Item{
			ID:          id,
			Name:        name,
			Description: description,
			Amount:      amount,
		}
		items = append(items, item)
	}

	return
}

// SelectItem - select 1 posts
func SelectItem(index int64) (item domain.Item, err error) {
	rows, err := Conn.Prepare("SELECT * FROM items WHERE id = ?")
	if err != nil {
		return
	}

	defer rows.Close()
	var id int64
	var name string
	var description string
	var amount int64
	err = rows.QueryRow(index).Scan(&id, &name, &description, &amount)
	if err != nil {
		return
	}

	item = domain.Item{
		ID:          id,
		Name:        name,
		Description: description,
		Amount:      amount,
	}

	return
}
