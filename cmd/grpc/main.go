package main

import (
	"net"

	"github.com/knipers/go-grpc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"database/sql"

	"github.com/knipers/go-grpc/internal/database"
	"github.com/knipers/go-grpc/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.sqlite")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	authorDB := database.NewAuthor(db)
	authorService := service.NewAuthorService(*authorDB)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthorServiceServer(grpcServer, authorService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
