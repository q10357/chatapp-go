package graph

type RequestParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

/*
func CatGraphRouter(c *gin.Context) {
	var reqObj RequestParams
	if err := c.ShouldBindJSON(&reqObj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         cat.CatRootSchema,
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
*/
