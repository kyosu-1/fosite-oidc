package handler

import (
    "crypto/rsa"
    "encoding/json"
    "net/http"

    "github.com/ory/fosite/token/jwt"
    "github.com/ory/fosite/token/jwt/rs256"
    jose "gopkg.in/square/go-jose.v2"
)

func JWKsHandler(signer *jwt.DefaultSigner) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        publicKey, ok := signer.GetPublicKey(context.Background()).(*rsa.PublicKey)
        if !ok {
            http.Error(w, "Failed to get public key", http.StatusInternalServerError)
            return
        }

        // JWKの生成
        jwk := jose.JSONWebKey{
            Key:       publicKey,
            Use:       "sig",
            Algorithm: string(rs256.RS256),
        }
        keySet := jose.JSONWebKeySet{
            Keys: []jose.JSONWebKey{jwk},
        }

        // JSONレスポンスの作成
        response, err := json.Marshal(keySet)
        if err != nil {
            http.Error(w, "Failed to marshal public keys", http.StatusInternalServerError)
            return
        }

        // レスポンスの書き込み
        w.Header().Set("Content-Type", "application/json")
        w.Write(response)
    }
}
