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
				Type: graphql.NewList(UserType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return models.GetUsers(), nil
				},
			},
			"user": &graphql.Field{
				Type:        UserType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						user := models.GetUser(id)
						return user, nil
					}
					return nil, nil
				},
			},
		},
	})
	return queryType
}
