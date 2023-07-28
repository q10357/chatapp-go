package middleware

import (
	"net/http"

	"github.com/q10357/AuthWGo/auth/jwt"
)

// later: use middleware to restrict users access to db (only their id)
// we want all our routes to be authenticated, here we validate the token
func TokenValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// is token present?
		//If not -> redirect to /signin (302)
		if _, ok := r.Header["token"]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token missing"))
			return
		}
		token := r.Header["token"][0]
		check, err := jwt.ValidateToken(token)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Token Validation Failed"))
			return
		}
		if !check {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token Invalid"))
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Authorized Token"))
		next.ServeHTTP(rw, r)
	})
}
