package schemas

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/q10357/RelService/services"
)

type Resolver struct {
	rs *services.RelService
}

func NewResolver(rs *services.RelService) *Resolver {
	return &Resolver{rs: rs}
}

var relType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Rel",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"otherUsername": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (r *Resolver) relsResolve(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	userIdValue := ctx.Value("userId")
	if userIdValue == nil {
		return nil, fmt.Errorf("userId not found in request headers")
	}

	userId, err := strconv.Atoi(userIdValue.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse userId from request headers")
	}

	return r.rs.GetRelsByUserId(uint(userId))
}

func (r *Resolver) CreateRelQueries() *graphql.Object {
	return graphql.NewObject(
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
					Resolve: r.relsResolve,
				},
			},
		},
	)
}

func NewRelRootSchema(rs *services.RelService) (*graphql.Schema, error) {
	resolver := NewResolver(rs)
	relQueries := resolver.CreateRelQueries()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: relQueries,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new schema, error: %v", err)
	}

	return &schema, nil
}
