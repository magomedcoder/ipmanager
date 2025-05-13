package server

import (
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AppProvider struct {
	UserHandler pb.UserServiceServer
	Middleware  middleware.AuthMiddleware
}

func Run(app *AppProvider) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("Ошибка при запуске слушателя: %v", err)
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(app.Middleware.UnaryInterceptor),
		grpc.StreamInterceptor(app.Middleware.StreamInterceptor),
	)

	pb.RegisterUserServiceServer(s, app.UserHandler)

	log.Printf("Сервер слушает на %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Printf("Ошибка при запуске сервера: %v", err)
		return err
	}

	return nil
}
