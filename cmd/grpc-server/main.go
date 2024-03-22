package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/elidelima/go-grpc/internal/database"
	"github.com/elidelima/go-grpc/internal/pb"
	"github.com/elidelima/go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	fmt.Print("Listening on TCP:50051")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
