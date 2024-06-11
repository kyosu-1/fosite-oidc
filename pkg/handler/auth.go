package handler

import (
    "context"
    "net/http"

    "github.com/ory/fosite"
)

func AuthHandler(oauth2 fosite.OAuth2Provider) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.Background()
        ar, err := oauth2.NewAuthorizeRequest(ctx, r)
        if err != nil {
            oauth2.WriteAuthorizeError(ctx, w, ar, err)
            return
        }

        // Simulate user login and consent
        username := r.URL.Query().Get("username")
        if username == "" {
            http.Redirect(w, r, "/login?challenge="+ar.GetID(), http.StatusFound)
            return
        }

        consent := r.URL.Query().Get("consent")
        if consent != "true" {
            oauth2.WriteAuthorizeError(ctx, w, ar, fosite.ErrConsentRequired)
            return
        }

        for _, scope := range ar.GetRequestedScopes() {
            ar.GrantScope(scope)
        }

        response, err := oauth2.NewAuthorizeResponse(ctx, ar, &fosite.DefaultSession{
            Subject: username,
        })
        if err != nil {
            oauth2.WriteAuthorizeError(ctx, w, ar, err)
            return
        }

        oauth2.WriteAuthorizeResponse(ctx, w, ar, response)
    }
}
