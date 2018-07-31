package handler

import (
	"net/http"
	"strconv"

	"vue-golang-payment-app/backend-api/db"
)

// GetList - get all items
func GetList(c Context) {
	res, err := db.SelectAllItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetItem - get one item
func GetItem(c Context) {
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, res)
}
