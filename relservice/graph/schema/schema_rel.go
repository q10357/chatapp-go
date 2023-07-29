package schema

import (
	"github.com/graphql-go/graphql"
)

var relType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Rel",
		Fields: graphql.Fields{
			"relId": &graphql.Field{
				Type: graphql.Int,
			},
			"user1": &graphql.Field{
				Type: userType,
			},
			"user2": &graphql.Field{
				Type: userType,
			},
		},
	},
)
