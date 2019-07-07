package graphs

import (
	"context"
	"errors"

	"github.com/YanickJair/AuthGo/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// NodeDefinitions - Define our node in here
var NodeDefinitions *relay.NodeDefinitions

// ShipType - type
var ShipType *graphql.Object

// FactionType - type
var FactionType *graphql.Object

// QueryType - object
var QueryType *graphql.Object

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

	// ShipType - define our ship object
	ShipType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Ship",
		Description: "A ship",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Ship", nil),
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the ship",
			},
		},
		Interfaces: []*graphql.Interface{
			NodeDefinitions.NodeInterface,
		},
	})

	// connection - define our connection
	var shipConnectionField = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "Ship",
		NodeType: ShipType,
	})

	// Faction definition
	FactionType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Faction",
		Description: "A faction",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Faction", nil),
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the ship",
			},
			"ships": &graphql.Field{
				Type: shipConnectionField.ConnectionType,
				Args: relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					ships := []interface{}{}
					if faction, ok := p.Source.(*models.Faction); ok {
						for _, shipID := range faction.Ships {
							ships = append(ships, models.GetShip(shipID))
						}
					}
					return relay.ConnectionFromArray(ships, args), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			NodeDefinitions.NodeInterface,
		},
	})

	// Define our query
	QueryType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Query data from source",
		Fields: graphql.Fields{
			"rebels": &graphql.Field{
				Type: FactionType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return models.GetRebels(), nil
				},
			},
			"empire": &graphql.Field{
				Type: FactionType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return models.GetEmpire(), nil
				},
			},
			"node": NodeDefinitions.NodeField,
		},
	})
}
