package handler

import (
    "context"
    "net/http"

    "github.com/ory/fosite"
)

func RevokeHandler(oauth2 fosite.OAuth2Provider) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.Background()
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        
        err := oauth2.NewRevocationRequest(ctx, r)
        if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}
