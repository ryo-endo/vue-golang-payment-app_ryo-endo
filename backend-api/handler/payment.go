package handler

import (
	"context"
	"net/http"
	"strconv"

	"vue-golang-payment-app/backend-api/db"
	"vue-golang-payment-app/backend-api/domain"
	gpay "vue-golang-payment-app/payment-service/proto"

	"google.golang.org/grpc"
)

func Charge(c Context) {
	t := domain.Payment{}
	c.Bind(&t)

	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// item取得
	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// gRPCサーバーに送る
	greq := &gpay.PayRequest{
		Id:          int64(identifer),
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	// gRPCと接続
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	defer conn.Close()

	client := gpay.NewPayManagerClient(conn)

	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	c.JSON(http.StatusOK, gres)

}
