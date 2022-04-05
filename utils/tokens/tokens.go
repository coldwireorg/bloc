package tokens

import (
	"bloc/utils/config"
	"time"

	"codeberg.org/coldwire/cwauth"
	"github.com/kataras/jwt"
	"github.com/rs/zerolog/log"
)

var key *string

type Token struct {
	Username   string `json:"username"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// Generate token:
//	In this function we just put the content of the token
//	like the username, the user id and the expiration time
//	of the token
func Generate(body Token, exp time.Duration) string {
	// Define the header
	header := jwt.Claims{
		Expiry:   time.Now().Add(exp).Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer:   "coldwire",
	}

	// Sign the token with a ed25519 private key
	t, err := jwt.Sign(jwt.HS256, []byte(*key), body, header)
	if err != nil {
		log.Warn().Msg(err.Error())
	}

	// return token as a string
	return string(t)
}

// Verify JWT tokens:
// 	This simply work by verifying the signature with
//	a ed25519 public key
func Verify(token string) (*jwt.VerifiedToken, error) {
	t := []byte(token) // get token

	// Verify token
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(*key), t)
	if err != nil {
		return nil, err
	}

	// return token
	return verifiedToken, nil
}

// This function is used for decoding JWT token
// WARNING: this is actually not verifying it!!
func Parse(token string) (Token, error) {
	// if oauth is used parse id_token
	if config.Conf.Oauth.Server != "" {
		claims := cwauth.GetClaims(token)

		return Token{
			Username:   claims.Username,
			PrivateKey: claims.PrivateKey,
			PublicKey:  claims.PublicKey,
		}, nil
	}

	t, err := jwt.Decode([]byte(token)) // get token from the cookie
	if err != nil {
		return Token{}, err // return void token payload with the error
	}

	// Get token payload content
	var payload Token
	err = t.Claims(&payload)
	if err != nil {
		return Token{}, err
	}

	// return the token payload without errors
	return payload, nil
}

func Init() {
	k, err := cwauth.GenerateRandomString(32)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	key = &k
}
