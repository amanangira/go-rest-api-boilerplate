package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/tozd/go/errors"
	"net/http"
	"strings"
	"superman/api"
	"superman/api/api_error"
	"superman/external/cognito"
)

const authorizationHeader = "Authorization"
const contextUserIDKey = "user_id"

func AuthorizedCognitoUserID(r *http.Request) (*cognito.Claims, error) {
	//token
	idTokenString := stripBearerToken(r.Header.Get(authorizationHeader))
	jwk, jwkErr := api.DefaultJWK()
	if jwkErr != nil {
		return nil, jwkErr
	}

	token, parseErr := jwk.ParseJWT(idTokenString)
	if parseErr != nil {
		return nil, parseErr
	}

	claims, claimParseErr := cognito.ParseClaims(token.Claims)
	if claimParseErr != nil {
		return nil, claimParseErr
	}

	return claims, nil
}

func stripBearerToken(token string) string {
	prefix := "bearer"
	if strings.HasPrefix(strings.ToLower(token), prefix) {
		parts := strings.Split(token, " ")
		if len(parts) == 2 {
			return parts[1]
		}
	}

	return ""
}

//TODO - Consider decoupling from IAPI and instead implement a middleware interface that can be implemented by a struct
// that would hold any dependencies for that particular middleware following isolation and decouple design

func ApplyBasicAuthorizer(api api.IAPI) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, cognitoUserIDErr := AuthorizedCognitoUserID(r)

			if claims == nil {
				api_error.NewError(cognitoUserIDErr, http.StatusUnauthorized, "").Log().HttpError(w)
				return
			}

			var count int
			dbUserID := claims.CognitoUsername
			queryErr := api.
				GetDBClient().
				GetContext(
					r.Context(),
					&count,
					`SELECT count(*) FROM users WHERE id = $1 LIMIT 1`,
					dbUserID)

			if queryErr != nil && queryErr != sql.ErrNoRows {
				api_error.NewError(queryErr, http.StatusInternalServerError, "").Log().HttpError(w)
				return
			}

			if queryErr == sql.ErrNoRows || count < 1 {
				api_error.NewError(errors.New("db user with Cognito username not found"), http.StatusUnauthorized, "").Log().HttpError(w)
				return
			}

			if cognitoUserIDErr != nil {
				api_error.NewError(cognitoUserIDErr, http.StatusUnauthorized, "").Log().HttpError(w)
				return
			}

			ctx := r.Context()
			if ctx.Value(contextUserIDKey) != dbUserID {
				ctxWithUserSession := context.WithValue(ctx, contextUserIDKey, dbUserID)
				r = r.WithContext(ctxWithUserSession)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func CurrentUserID(ctx context.Context) string {
	userID := fmt.Sprint(ctx.Value(contextUserIDKey))

	return userID
}
