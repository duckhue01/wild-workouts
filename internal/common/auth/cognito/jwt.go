package cognito

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tribefintech/microservices/internal/common/cmerr"
)

const (
	KEY_NOT_FOUND_ERR = "Key not found"
)

// Auth struct
type Auth struct {
	jwk               *JWK
	jwkURL            string
	cognitoRegion     string
	cognitoUserPoolID string
}

// JWK struct
type JWK struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
	} `json:"keys"`
}

func New(region, poolId string) *Auth {
	a := &Auth{
		cognitoRegion:     region,
		cognitoUserPoolID: poolId,
		jwkURL:            fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, poolId),
	}

	err := a.getJWK()
	if err != nil {
		panic(err)
	}

	return a
}

func (a *Auth) getJWK() error {
	req, err := http.NewRequest("GET", a.jwkURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	jwk := new(JWK)
	err = json.Unmarshal(body, jwk)
	if err != nil {
		return err
	}

	a.jwk = jwk
	return nil
}

// Parse parse JWT token into Token struct
func (a *Auth) Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		index := -1
		for i, v := range a.jwk.Keys {
			if v.Kid == token.Header["kid"] {
				index = i
			}
		}
		if index == -1 {
			return nil, cmerr.NewUnexpectedError(
				KEY_NOT_FOUND_ERR,
				cmerr.InternalServerError)
		}
		if token.Method.Alg() != "RS256" {

			return nil, cmerr.NewAuthorizationError(
				fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]),
				cmerr.UnexpectedSigningMethod)
		}
		key := convertKey(a.jwk.Keys[index].E, a.jwk.Keys[index].N)
		return key, nil
	})

	return token, err
}

// JWK return JWK
func (a *Auth) JWK() *JWK {
	return a.jwk
}

// JWKURL return JWK URL
func (a *Auth) JWKURL() string {
	return a.jwkURL
}

func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey
}
