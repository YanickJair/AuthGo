package graphs

import (
	"github.com/graphql-go/graphql"
)

// MutationType - func mutation
func MutationType() *graphql.Object {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"signUp": UserMutation(),
		},
	})

	return mutationType
}
