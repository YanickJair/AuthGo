package graphs

/*
import (
	"context"
	"errors"

	"github.com/YanickJair/AuthGo/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// NodeDefinitions - Define our node in here
var NodeDefinitions *relay.NodeDefinitions

func init() {
	NodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo,
			ctx context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			switch resolvedID.Type {
			case "Faction":
				return models.GetFaction(resolvedID.ID), nil
			case "Ship":
				return models.GetShip(resolvedID.ID), nil
			default:
				return nil, errors.New("Unkown type")
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *models.Faction:
				return FactionType
			default:
				return ShipType
			}
		},
	})
}
*/
