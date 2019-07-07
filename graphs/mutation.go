package graphs

import (
	"context"

	"github.com/YanickJair/AuthGo/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var shipMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "IntroduceShip",
	InputFields: graphql.InputObjectConfigFieldMap{
		"shipName": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"factionId": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	OutputFields: graphql.Fields{
		"ship": &graphql.Field{
			Type: ShipType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return models.GetShip(payload["shipId"].(string)), nil
				}
				return nil, nil
			},
		},
		"faction": &graphql.Field{
			Type: FactionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return models.GetFaction(payload["factionId"].(string)), nil
				}
				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		// `inputMap` is a map with keys/fields as specified in `InputFields`
		// Note, that these fields were specified as non-nullables, so we can assume that it exists.
		shipName := inputMap["shipName"].(string)
		factionID := inputMap["factionId"].(string)

		// This mutation involves us creating (introducing) a new ship
		newShip := models.CreateShip(shipName, factionID)
		// return payload
		return map[string]interface{}{
			"shipId":    newShip.ID,
			"factionId": factionID,
		}, nil
	},
})

// MutationType - mutation
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"introduceShip": shipMutation,
	},
})
