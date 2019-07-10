package main

import (
	"fmt"
	"net/http"

	"github.com/YanickJair/AuthGo/graphs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphs.QueryType(),
		Mutation: graphs.MutationType(),
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	fs := http.FileServer(http.Dir("static"))

	fmt.Println("Server running: http://localhost:8080")

	http.Handle("/graphql", h)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
