package schemas

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/q10357/RelService/services"
)

type RelResolver struct {
	rs *services.RelService
}

func NewRelResolver(rs *services.RelService) *RelResolver {
	return &RelResolver{rs: rs}
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
			"isRequester": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

func (r *RelResolver) getUserIdFromContext(p *graphql.ResolveParams) (*uint, error) {
	// Get the userID from the context and assert it to an int
	userIDStr := p.Context.Value("userId").(string)

	if userIDStr == "" {
		return nil, fmt.Errorf("userID not found in context")
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)

	if err != nil {
		return nil, fmt.Errorf("error converting ID to int")
	}

	userIDUint := uint(userID)
	return &userIDUint, nil
}

func (r *RelResolver) isUserPartOfRelationship(userID, relId uint) bool {
	userIsPartOfRel, err := r.rs.IsUserIsInRelation(relId, userID)

	if err != nil {
		return false
	}

	if !userIsPartOfRel {
		return false
	}

	return true

}

func (r *RelResolver) CreateRelQueries() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RelQueries",
			Fields: graphql.Fields{
				"rels": &graphql.Field{
					Type: graphql.NewList(relType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						userID, err := r.getUserIdFromContext(&p)
						if err != nil {
							return nil, err
						}
						return r.rs.GetRelsByUserId(uint(*userID))
					},
				},
			},
		},
	)
}

func (r *RelResolver) CreateRelMutations() *graphql.Object {
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
	resolver := NewRelResolver(rs)
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
