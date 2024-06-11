package server

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ory/fosite"
    "github.com/ory/fosite/compose"
	"github.com/kyosu-1/fosite-oidc/pkg/handler"
    "github.com/ory/fosite/storage"
    "github.com/ory/fosite/token/jwt"
    "context"
    "time"
)

var (
    store fosite.Storage
    oauth2 fosite.OAuth2Provider
    signer *jwt.DefaultSigner
)

func init() {
    store = storage.NewMemoryStore()
    config := &compose.Config{
        AccessTokenLifespan: time.Hour,
    }

    // 秘密鍵と公開鍵の読み込み
    privateKey, publicKey := loadKeys("cert/private.pem", "cert/public.pem")

    signer = &jwt.DefaultSigner{
        GetPrivateKey: func(_ context.Context) (interface{}, error) {
            return privateKey, nil
        },
        GetPublicKey: func(_ context.Context) (interface{}, error) {
            return publicKey, nil
        },
    }

    strategy := compose.CommonStrategy{
        CoreStrategy: compose.NewOAuth2HMACStrategy(config, []byte("some-super-cool-secret-that-nobody-knows")),
        OpenIDStrategy: compose.NewOpenIDConnectStrategy(config, signer),
    }

    oauth2 = compose.Compose(
        config,
        store,
        strategy,
        compose.OAuth2AuthorizeExplicitFactory,
        compose.OAuth2TokenIntrospectionFactory,
        compose.OAuth2RefreshTokenGrantFactory,
        compose.OpenIDConnectExplicitFactory,
    )
}

func loadKeys(privatePath, publicPath string) (*rsa.PrivateKey, *rsa.PublicKey) {
    // 秘密鍵の読み込み
    privateKeyData, err := ioutil.ReadFile(privatePath)
    if err != nil {
        log.Fatalf("Failed to read private key: %v", err)
    }
    block, _ := pem.Decode(privateKeyData)
    if block == nil || block.Type != "RSA PRIVATE KEY" {
        log.Fatalf("Failed to decode PEM block containing private key")
    }
    privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }

    // 公開鍵の読み込み
    publicKeyData, err := ioutil.ReadFile(publicPath)
    if err != nil {
        log.Fatalf("Failed to read public key: %v", err)
    }
    block, _ = pem.Decode(publicKeyData)
    if block == nil || block.Type != "PUBLIC KEY" {
        log.Fatalf("Failed to decode PEM block containing public key")
    }
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        log.Fatalf("Failed to parse public key: %v", err)
    }
    publicKey, ok := pub.(*rsa.PublicKey)
    if !ok {
        log.Fatalf("Failed to cast public key to RSA public key")
    }

    return privateKey, publicKey
}

func Start() {
    r := mux.NewRouter()
    handler.RegisterHandlers(r, oauth2, store)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
