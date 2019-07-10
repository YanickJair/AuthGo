package graphs

import (
	"context"
	"errors"

	"github.com/YanickJair/AuthGo/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// NodeDefinitions - fun
func NodeDefinitions() *relay.NodeDefinitions {
	nodeDefinitions := relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo,
			ctx context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			switch resolvedID.Type {
			case "User":
				return models.GetUsers(), nil
			default:
				return nil, errors.New("Error")
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *models.User:
				return UserType
			default:
				return UserType
			}
		},
	})
	return nodeDefinitions
}
