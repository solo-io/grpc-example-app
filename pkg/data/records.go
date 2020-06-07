package data

import "github.com/solo-io/grpc-example-app/api/store"

var Records = []*store.Record{
	{
		Uuid:   "4",
		Name:   "OK Computer",
		Artist: "Radiohead",
		Price:  24.20,
	},
	{
		Uuid:   "5",
		Name:   "Nevermind",
		Artist: "Nirvana",
		Price:  28.69,
	},
	{
		Uuid:   "6",
		Name:   "Sublime",
		Artist: "Sublime",
		Price:  40.45,
	},
}
