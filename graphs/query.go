package graphs

import (
	"github.com/YanickJair/AuthGo/models"
	"github.com/graphql-go/graphql"
)

// QueryType - func query
func QueryType() *graphql.Object {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return models.GetUsers(), nil
				},
			},
		},
	})
	return queryType
}
