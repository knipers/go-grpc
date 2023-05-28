package service

import (
	"context"
	"io"

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

func (a *AuthorService) FindById(ctx context.Context, in *pb.AuthorGetRequest) (*pb.Author, error) {
	author, err := a.AuthorDB.FindById(in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Author{
		Id:   author.ID,
		Name: author.Name,
	}, nil
}

func (a *AuthorService) CreateAuthorStream(stream pb.AuthorService_CreateAuthorStreamServer) error {
	authors := &pb.AuthorList{}

	for {
		author, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(authors)
		}

		if err != nil {
			return err
		}

		authorResult, err := a.AuthorDB.Create(author.Name)

		if err != nil {
			return err
		}

		authors.Authors = append(authors.Authors, &pb.Author{
			Id:   authorResult.ID,
			Name: authorResult.Name,
		})
	}
}

func (a *AuthorService) CreateAuthorBidirectional(stream pb.AuthorService_CreateAuthorBidirectionalServer) error {
	for {
		author, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		authorResult, err := a.AuthorDB.Create(author.Name)

		if err != nil {
			return err
		}

		if err := stream.Send(&pb.Author{
			Id:   authorResult.ID,
			Name: authorResult.Name,
		}); err != nil {
			return err
		}
	}

}
