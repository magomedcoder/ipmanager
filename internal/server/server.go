package server

import (
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/handler"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net/http"
	"strings"
)

type AppProvider struct {
	UserHandler pb.UserServiceServer
	Middleware  middleware.AuthMiddleware
}

func Run(app *AppProvider) error {
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(app.Middleware.UnaryInterceptor),
		grpc.StreamInterceptor(app.Middleware.StreamInterceptor),
	)

	pb.RegisterUserServiceServer(srv, app.UserHandler)

	reflection.Register(srv)

	//listener, err := net.Listen("tcp", ":50051")
	//if err != nil {
	//	log.Printf("Ошибка при запуске слушателя: %v", err)
	//	return err
	//}
	//
	//go func() {
	//	if err := srv.Serve(listener); err != nil {
	//		log.Fatalf("Ошибка работы gRPC-сервера: %v", err)
	//		return
	//	}
	//}()

	wrappedGrpc := grpcweb.WrapServer(srv, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))

	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(r) || wrappedGrpc.IsAcceptableGrpcCorsRequest(r) || wrappedGrpc.IsGrpcWebSocketRequest(r) {
			wrappedGrpc.ServeHTTP(w, r)
			return
		}

		fileServer := http.FileServer(handler.UIStaticFiles)

		if strings.HasPrefix(r.URL.Path, "/static/") {
			fileServer.ServeHTTP(w, r)
			return
		}

		indexFile, err := handler.UIStaticFiles.Open("index.html")
		if err != nil {
			http.Error(w, "Не удалось загрузить index.html", http.StatusInternalServerError)
			return
		}
		defer func(indexFile http.File) {
			if err := indexFile.Close(); err != nil {
				fmt.Println(err)
			}
		}(indexFile)

		data, err := io.ReadAll(indexFile)
		if err != nil {
			http.Error(w, "Ошибка чтения index.html", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write(data); err != nil {
			fmt.Println(err)
			return
		}
	})

	httpServer := &http.Server{
		Addr:    ":8000",
		Handler: mainHandler,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка HTTP-сервера: %v", err)
	}

	//log.Printf("Сервер слушает на %v", listener.Addr())
	//
	//if err := srv.Serve(listener); err != nil {
	//	log.Printf("Ошибка при запуске сервера: %v", err)
	//	return err
	//}

	return nil
}
