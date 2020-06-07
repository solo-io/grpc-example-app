package data

import "github.com/solo-io/grpc-example-app/api/store"

var Books = []*store.Book{
	{
		Uuid:                 "1",
		Name:                 "To Kill a Mockingbird",
		Author:               "Harper Lee",
		Price:                7.19,
	},
	{
		Uuid:                 "2",
		Name:                 "Catcher in the Rye",
		Author:               "J.D. Salinger",
		Price:                5.99,
	},
	{
		Uuid:                 "3",
		Name:                 "Lord of the Flies",
		Author:               "William Golding",
		Price:                5.99,
	},
}