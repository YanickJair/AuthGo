package main

import (
	"fmt"
	"net/http"

	"github.com/YanickJair/AuthGo/graphs"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &graphs.Schema,
		Pretty: true,
	})

	fs := http.FileServer(http.Dir("static"))

	fmt.Println("Server running: http://localhost:8080")

	http.Handle("/graphql", h)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
