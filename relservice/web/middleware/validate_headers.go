package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDString := c.GetHeader("userId")
		if userIDString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "UserID not found in header"})
			c.Abort()
			return
		}

		// Add the userId to the context
		c.Set("userId", userIDString)

		// Process the request
		c.Next()
	}
}

func checkUserInRelDatabase(userID uint) bool {
	// Query the relationship database to check if the user is in a relationship.
	// This is just a placeholder, replace with your actual database query.

	// For example:
	// db, err := sql.Open("mysql", "user:password@/dbname")
	// if err != nil { ... handle this error ... }
	// defer db.Close()

	// rows, err := db.Query("SELECT * FROM relationships WHERE userIdRequester = ? OR userIdRequested = ?", userID, userID)
	// if err != nil { ... handle this error ... }
	// defer rows.Close()

	// For simplicity, let's assume the user is always found in the database
	return true
}
