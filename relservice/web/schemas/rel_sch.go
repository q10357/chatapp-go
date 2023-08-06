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

type ResolverFunc func(p graphql.ResolveParams) (interface{}, error)

func (r *Resolver) resolveAuthHeader(p *graphql.ResolveParams) (*int, error) {
	ctx := p.Context

	userIdValue := ctx.Value("userId")
	if userIdValue == nil {
		return nil, fmt.Errorf("userId not found in request headers")
	}

	userId, err := strconv.Atoi(userIdValue.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse userId from request headers")
	}

	return &userId, nil
}

func (r *Resolver) isUserPartOfRelationship(userId, relId uint) bool {
	// Implement your logic to check the relationship.
	// You will probably query your relationship database here.
	// For simplicity, I will return true.
	userIsPartOfRel, err := r.rs.IsUserIsInRelation(relId, userId)

	if err != nil {
		return false
	}

	if !userIsPartOfRel {
		return false
	}

	return true

}

func (r *Resolver) CreateRelQueries() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RelQueries",
			Fields: graphql.Fields{
				"rels": &graphql.Field{
					Type: graphql.NewList(relType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						userId, err := r.resolveAuthHeader(&p)
						if err != nil {
							return nil, err
						}
						// Rest of your resolver logic goes here
						return r.rs.GetRelsByUserId(uint(*userId))
					},
				},
			},
		},
	)
}

func (r *Resolver) CreateRelMutations() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RelMutations",
			Fields: graphql.Fields{
				"acceptRel": &graphql.Field{
					Type: relType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					/*Here i will accept the incitation from the requester, how can i check the user is a part of relationship here also?*/
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
