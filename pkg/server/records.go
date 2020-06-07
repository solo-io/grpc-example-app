package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/solo-io/grpc-example-app/api/store"
)

type recordServer struct {
	records []*store.Record
}

func NewRecordServer(records []*store.Record) *recordServer {
	return &recordServer{records: records}
}

func (b *recordServer) ListRecords(context.Context, *empty.Empty) (*store.ListRecordsResponse, error) {
	return &store.ListRecordsResponse{
		Records: b.records,
	}, nil
}
