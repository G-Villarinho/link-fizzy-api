package middlewares

import (
	"net/http"
	"strings"

	"github.com/g-villarinho/link-fizz-api/pkgs/di"
	"github.com/g-villarinho/link-fizz-api/pkgs/requestcontext"
	"github.com/g-villarinho/link-fizz-api/responses"
	"github.com/g-villarinho/link-fizz-api/services"
)

type errorResponse struct {
	Code string `json:"code"`
}

type AuthMiddleware interface {
	Authenticate(next http.Handler) http.Handler
}

type authMiddleware struct {
	i  *di.Injector
	ts services.TokenService
	ls services.LogoutService
	rc requestcontext.RequestContext
}

func NewAuthMiddleware(i *di.Injector) (AuthMiddleware, error) {
	tokenService, err := di.Invoke[services.TokenService](i)
	if err != nil {
		return nil, err
	}

	logoutService, err := di.Invoke[services.LogoutService](i)
	if err != nil {
		return nil, err
	}

	requestContext, err := di.Invoke[requestcontext.RequestContext](i)
	if err != nil {
		return nil, err
	}

	return &authMiddleware{
		i:  i,
		ts: tokenService,
		ls: logoutService,
		rc: requestContext,
	}, nil
}

func (a *authMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			responses.JSON(w, http.StatusUnauthorized, errorResponse{
				Code: "UNAUTHORIZED",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			responses.JSON(w, http.StatusUnauthorized, errorResponse{
				Code: "UNAUTHORIZED",
			})
			return
		}

		token := parts[1]

		claims, err := a.ts.ValidateToken(r.Context(), token)
		if err != nil {
			responses.NoContent(w, http.StatusUnauthorized)
			return
		}

		logout, err := a.ls.IsLogoutRegistered(r.Context(), token)
		if err != nil {
			responses.JSON(w, http.StatusUnauthorized, errorResponse{
				Code: "UNAUTHORIZED",
			})
			return
		}

		if logout {
			responses.JSON(w, http.StatusUnauthorized, errorResponse{
				Code: "UNAUTHORIZED",
			})
			return
		}

		ctx := a.rc.SetUserID(r.Context(), claims.Sub)
		ctx = a.rc.SetToken(ctx, token)
		ctx = a.rc.SetSessionID(ctx, claims.Sid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
