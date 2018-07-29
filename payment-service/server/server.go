package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gpay "../proto"
	payjp "github.com/payjp/payjp-go/v1"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Charge(ctx context.Context, req *gpay.PayRequest) (*gpay.PayResponse, error) {
	// 初期化
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY"), nil)

	// 支払いの実行
	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token,
		Capture:     true,
		Description: req.Name + ":" + req.Description,
	})
	if err != nil {
		return nil, err
	}

	// 支払い結果からResponse生成
	res := &gpay.PayResponse{
		Paid:     charge.Paid,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	s := grpc.NewServer()
	gpay.RegisterPayManagerServer(s, &server{})

	reflection.Register(s)
	log.Printf("gRPC server started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
