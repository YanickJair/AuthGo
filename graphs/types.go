package graphs

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// UserType - user
var UserType *graphql.Object

// NodeDefinitions - node
var nodeDefinitions = NodeDefinitions()

func init() {
	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User type",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"fname": &graphql.Field{
				Type:        graphql.String,
				Description: "First name",
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "Email address",
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})
}
