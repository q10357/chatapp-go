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

func NewRelGraphRouter(schema *graphql.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj RequestParams

		// Check if the userId was set in the middleware
		userIdStr, exists := c.Get("userId")

		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "userId not found in context"})
			return
		}

		// No need to set the userId again in the context, it was already set in the middleware

		if err := c.ShouldBindJSON(&reqObj); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a new context with userId
		ctx := context.WithValue(c.Request.Context(), "userId", userIdStr)

		result := graphql.Do(graphql.Params{
			Context:        ctx, // Pass the existing Gin context, which already includes the userId
			Schema:         *schema,
			RequestString:  reqObj.Query,
			VariableValues: reqObj.Variables,
			OperationName:  reqObj.Operation,
		})

		if len(result.Errors) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Errors})
			return
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
