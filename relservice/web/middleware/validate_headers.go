package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func validateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDString := c.GetHeader("UserID")
		if userIDString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "UserID not found in header"})
			c.Abort()
			return
		}

		userID, err := strconv.ParseUint(userIDString, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid UserID"})
			c.Abort()
			return
		}

		if !checkUserInRelDatabase(uint(userID)) {
			c.JSON(http.StatusForbidden, gin.H{"status": "User not found in the relationship database"})
			c.Abort()
			return
		}

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
