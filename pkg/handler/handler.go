package handler

import (
    "github.com/gorilla/mux"
    "github.com/ory/fosite"
    "github.com/ory/fosite/token/jwt"
)

func RegisterHandlers(r *mux.Router, oauth2 fosite.OAuth2Provider, store fosite.Storage) {
    r.HandleFunc("/oauth2/auth", AuthHandler(oauth2)).Methods("GET", "POST")
    r.HandleFunc("/oauth2/token", TokenHandler(oauth2)).Methods("POST")
    r.HandleFunc("/oauth2/revoke", RevokeHandler(oauth2)).Methods("POST")
    r.HandleFunc("/oauth2/introspect", IntrospectHandler(oauth2)).Methods("POST")
    r.HandleFunc("/.well-known/jwks.json", JWKsHandler(&jwt.DefaultSigner{})).Methods("GET")
    r.HandleFunc("/.well-known/openid-configuration", OIDCConfigurationHandler()).Methods("GET")
}
