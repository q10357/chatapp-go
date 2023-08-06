package graph

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type RequestParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// NewRelGraphRouter initializes a new GraphQL route handler with the provided schema.
func NewRelGraphRouter(schema *graphql.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj RequestParams

		// Extract userId from context
		userIdStr, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found in context"})
			return
		}

		// Bind request data to reqObj
		if err := c.ShouldBindJSON(&reqObj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a new context with userId
		ctx := context.WithValue(c.Request.Context(), "userId", userIdStr)

		// Execute GraphQL query
		result := graphql.Do(graphql.Params{
			Context:        ctx,
			Schema:         *schema,
			RequestString:  reqObj.Query,
			VariableValues: reqObj.Variables,
			OperationName:  reqObj.Operation,
		})

		// Handle GraphQL execution errors
		if len(result.Errors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Errors})
			return
		}

		// Respond with successful result
		c.JSON(http.StatusOK, result)
	}
}
