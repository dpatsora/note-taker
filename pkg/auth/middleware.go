package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/dpatsora/note-taker/pkg/server/httperr"
)

func JWTMiddleware(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var claims jwt.MapClaims

			token, err := request.ParseFromRequest(
				r,
				request.AuthorizationHeaderExtractor,
				func(token *jwt.Token) (i interface{}, e error) {
					return []byte(secret), nil
				},
				request.WithClaims(&claims),
			)
			if err != nil {
				httperr.Unauthorised("unable-to-get-jwt", err, w, r)
				return
			}

			if !token.Valid {
				httperr.Unauthorised("invalid-jwt", nil, w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

}
