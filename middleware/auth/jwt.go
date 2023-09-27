package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aminkhn/portfolio/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidate = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("token signature is invalid")
)

func NewJWT() (*JWT, error) {
	env, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	return &JWT{
		[]byte(env.Secret),
	}, nil
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) (CustomClaims, error) {
	env, err := config.LoadConfig()
	if err != nil {
		return CustomClaims{}, err
	}
	bf, _ := ParseDuration(env.JBT)
	ep, _ := ParseDuration(env.JET)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    env.Issuer,
		},
	}
	return claims, nil
}
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	// Define your secret key used for signing the token
	secretKey := []byte(j.SigningKey)

	// Define a function to provide the key for validation
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// handle error
	}

	// Check if the token is valid
	if !token.Valid {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, nil, ErrTokenMalformed
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, nil, ErrTokenInvalid
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, nil, ErrTokenExpired
		} else {
			return nil, nil, fmt.Errorf("token is invalid: %w", err)
		}
	}

	// Return the validated token
	return token, claims, nil
}
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
