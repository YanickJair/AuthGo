package graphs

import (
	"github.com/graphql-go/graphql"
)

// Schema - main schema endpoint
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationType,
})
