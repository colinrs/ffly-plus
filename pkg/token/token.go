package token

import (
	"context"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `token` header was empty.
	ErrMissingHeader = errors.New("the length of the `token` header is zero")
)

// Context is the context of the JSON web token.
type Context struct {
	UserID         uint64
	Username       string
	ExpirationTime int64
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.UserID = uint64(claims["user_id"].(float64))
		ctx.Username = claims["username"].(string)
		ctx.ExpirationTime = int64(claims["expiration_time"].(float64))
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	token := c.Request.Header.Get("token")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(token) == 0 {
		return &Context{}, ErrMissingHeader
	}

	return Parse(token, secret)
}

// Sign signs the context with the specified secret.
func Sign(ctx context.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// The token content.
	// jti: （JWT ID）用于标识JWT的唯一ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         c.UserID,
		"username":        c.Username,
		"expiration_time": time.Now().Unix() + 24*3600*c.ExpirationTime,
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
