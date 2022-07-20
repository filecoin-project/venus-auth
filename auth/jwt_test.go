package auth

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/filecoin-project/venus-auth/config"
	"github.com/filecoin-project/venus-auth/storage"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	secret, err := config.RandSecret()
	if err != nil {
		panic(err)
	}
	sec, err := hex.DecodeString(hex.EncodeToString(secret))
	if err != nil {
		panic(err)
	}
	jwtOAuthInstance = &jwtOAuth{
		secret: jwt.NewHS256(sec),
		store:  storage.MockStore{},
		mp:     newMapper(),
	}
}

func TestGetToken(t *testing.T) {
	tokenInfo, err := jwtOAuthInstance.GetToken(context.Background(), storage.MockToken)
	assert.Nil(t, err)
	assert.Equal(t, string(storage.MockKeyPair.Token), tokenInfo.Token)
	assert.Equal(t, storage.MockKeyPair.Perm, tokenInfo.Perm)
}

// TODO: More tests

func TestTokenDecode(t *testing.T) {
	payload := []byte("eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ")
	pb, err := DecodeToBytes(payload)
	if err != nil {
		t.Fatal(err)
	}
	a := map[string]interface{}{}
	err = json.Unmarshal(pb, &a)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, a["name"], "John Doe")
}
