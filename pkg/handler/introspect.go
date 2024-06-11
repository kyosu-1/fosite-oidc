package handler

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/ory/fosite"
)

func IntrospectHandler(oauth2 fosite.OAuth2Provider) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.Background()
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        
        introspection, err := oauth2.NewIntrospectionRequest(ctx, r, &fosite.DefaultSession{})
        if err != nil {
            oauth2.WriteIntrospectionError(ctx, w, err)
            return
        }

        json.NewEncoder(w).Encode(introspection)
    }
}
