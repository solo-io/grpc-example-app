package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/solo-io/grpc-example-app/api/store"
)

type bookServer struct {
	books []*store.Book
}

func NewBookServer(books []*store.Book) *bookServer {
	return &bookServer{books: books}
}

func (b *bookServer) ListBooks(context.Context, *empty.Empty) (*store.ListBooksResponse, error) {
	return &store.ListBooksResponse{
		Books:                b.books,
	}, nil
}


