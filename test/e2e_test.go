package test_test

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/grpc-example-app/api/store"
	"github.com/solo-io/grpc-example-app/pkg/data"
	"google.golang.org/grpc"
	"os"
	"time"
)

var _ = Describe("Test", func() {

	It("can call the books microservice", func() {
		booksAddr := os.Getenv("BOOKS_ADDR")
		if booksAddr == "" {
			Skip("set BOOKS_ADDR to test the books service")
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		cc, err := grpc.DialContext(ctx, booksAddr, grpc.WithBlock(), grpc.WithInsecure())
		Expect(err).NotTo(HaveOccurred())

		bookServiceClient := store.NewBooksClient(cc)

		ctx = context.Background()

		resp, err := bookServiceClient.ListBooks(ctx, &empty.Empty{})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Books).To(Equal(data.Books))
	})

	It("can call the records microservice", func() {
		recordsAddr := os.Getenv("RECORDS_ADDR")
		if recordsAddr == "" {
			Skip("set RECORDS_ADDR to test the records service")
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		cc, err := grpc.DialContext(ctx, recordsAddr, grpc.WithBlock(), grpc.WithInsecure())
		Expect(err).NotTo(HaveOccurred())

		recordServiceClient := store.NewRecordsClient(cc)

		ctx = context.Background()

		resp, err := recordServiceClient.ListRecords(ctx, &empty.Empty{})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Records).To(Equal(data.Records))
	})
})
