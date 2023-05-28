package service

import (
	"context"

	"github.com/knipers/go-grpc/internal/database"
	"github.com/knipers/go-grpc/internal/pb"
)

type AuthorService struct {
	pb.UnimplementedAuthorServiceServer
	AuthorDB database.Author
}

func NewAuthorService(authorDB database.Author) *AuthorService {
	return &AuthorService{
		AuthorDB: authorDB,
	}
}

func (a *AuthorService) CreateAuthor(ctx context.Context, in *pb.CreateAuthorRequest) (*pb.Author, error) {
	author, err := a.AuthorDB.Create(in.Name)

	if err != nil {
		return nil, err
	}

	return &pb.Author{
		Id:   author.ID,
		Name: author.Name,
	}, nil
}

func (a *AuthorService) ListAuthors(ctx context.Context, _ *pb.Blank) (*pb.AuthorList, error) {
	authors, err := a.AuthorDB.FindAll()

	if err != nil {
		return nil, err
	}

	var pbAuthors []*pb.Author

	for _, author := range authors {
		pbAuthors = append(pbAuthors, &pb.Author{
			Id:   author.ID,
			Name: author.Name,
		})
	}

	return &pb.AuthorList{
		Authors: pbAuthors,
	}, nil

}
