package handler

import (
    "encoding/json"
    "net/http"
)

type OIDCConfiguration struct {
    Issuer   string `json:"issuer"`
    JwksURI  string `json:"jwks_uri"`
    AuthURL  string `json:"authorization_endpoint"`
    TokenURL string `json:"token_endpoint"`
    UserInfoURL string `json:"userinfo_endpoint"`
}

func OIDCConfigurationHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        config := OIDCConfiguration{
            Issuer:   "http://localhost:8080",
            JwksURI:  "http://localhost:8080/.well-known/jwks.json",
            AuthURL:  "http://localhost:8080/oauth2/auth",
            TokenURL: "http://localhost:8080/oauth2/token",
            UserInfoURL: "http://localhost:8080/userinfo",
        }

        json.NewEncoder(w).Encode(config)
    }
}
