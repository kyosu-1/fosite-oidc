package handler

import (
    "context"
    "net/http"

    "github.com/ory/fosite"
)

func TokenHandler(oauth2 fosite.OAuth2Provider) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.Background()
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        
        requester, err := oauth2.NewAccessRequest(ctx, r, &fosite.DefaultSession{})
        if err != nil {
            oauth2.WriteAccessError(ctx, w, requester, err)
            return
        }

        response, err := oauth2.NewAccessResponse(ctx, requester)
        if err != nil {
            oauth2.WriteAccessError(ctx, w, requester, err)
            return
        }

        oauth2.WriteAccessResponse(ctx, w, requester, response)
    }
}
