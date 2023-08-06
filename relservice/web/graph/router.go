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

		userId := c.GetHeader("userId")
		ctx := context.WithValue(c.Request.Context(), "userId", userId)
		c.Request = c.Request.WithContext(ctx)

		if err := c.ShouldBindJSON(&reqObj); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := graphql.Do(graphql.Params{
			Context:        ctx,
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
