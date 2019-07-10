package graphs

import (
	"context"

	"github.com/YanickJair/AuthGo/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// UserMutation - mutate user
func UserMutation() *graphql.Field {
	userMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "SignUp",
		InputFields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "your email address",
			},
			"fname": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "your first name",
			},
		},
		OutputFields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if payload, ok := p.Source.(map[string]interface{}); ok {
						return models.GetUser(payload["id"].(string)), nil
					}
					return nil, nil
				},
			},
		},
		MutateAndGetPayload: func(inputMap map[string]interface{},
			info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			email := inputMap["email"].(string)
			fname := inputMap["fname"].(string)
			newUser := models.SignUp(email, fname)
			return map[string]interface{}{
				"id": newUser.ID,
			}, nil
		},
	})
	return userMutation
}
