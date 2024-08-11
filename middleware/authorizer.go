package middleware

import (
	"boilerplate/app"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"boilerplate/app/app_error"

	"gitlab.com/tozd/go/errors"
)

const authorizationHeader = "Authorization"
const contextUserIDKey = "user_id"

func AuthorizedCognitoUserID(r *http.Request) (string, error) {
	//token
	var userID string
	idTokenString := stripBearerToken(r.Header.Get(authorizationHeader))
	// TODO - parse and return user ID from authorization header

	// TODO - temporarily hold token and user ID in DB
	userID = idTokenString

	return userID, nil
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

func ApplyBasicAuthorizer() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbUserID, cognitoUserIDErr := AuthorizedCognitoUserID(r)

			if cognitoUserIDErr != nil {
				app_error.NewError(cognitoUserIDErr, http.StatusUnauthorized, "").Log().HttpError(w)
				return
			}

			var count int
			queryErr := app.
				GetDBClient().
				GetContext(
					r.Context(),
					&count,
					`SELECT count(*) FROM users WHERE id = $1 LIMIT 1`,
					dbUserID)

			if queryErr != nil && queryErr != sql.ErrNoRows {
				app_error.NewError(queryErr, http.StatusInternalServerError, "").Log().HttpError(w)
				return
			}

			if queryErr == sql.ErrNoRows || count < 1 {
				app_error.NewError(errors.New("db user with Cognito username not found"), http.StatusUnauthorized, "").Log().HttpError(w)
				return
			}

			if cognitoUserIDErr != nil {
				app_error.NewError(cognitoUserIDErr, http.StatusUnauthorized, "").Log().HttpError(w)
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
