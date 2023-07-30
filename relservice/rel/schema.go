package rel

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var relType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Rel",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user1": &graphql.Field{
				Type: graphql.Int,
			},
			"user2": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var relQueries = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RelQueries",
		Fields: graphql.Fields{
			"rels": &graphql.Field{
				Type: graphql.NewList(relType),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ctx := p.Context

					userIdValue := ctx.Value("userId")
					if userIdValue == nil {
						return nil, fmt.Errorf("userId not found in request headers")
					}

					userId, err := strconv.Atoi(userIdValue.(string))
					if err != nil {
						return nil, fmt.Errorf("failed to parse userId from request headers")
					}

					return GetRelsByUserId(userId), nil
				},
			},
		},
	},
)

// used for root Schema
var RelRootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: relQueries,
	},
)
