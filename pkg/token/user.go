package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/xybor/xyauth/internal/models"
	"github.com/xybor/xyauth/internal/utils"
)

type AccessTokenClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}

var privateKey []byte
var publicKey []byte
var accessTokenExp = utils.GetConfig().GetDefault(
	"oauth.access_token_exp", 5*time.Minute).MustDuration()

func init() {
	var err error
	privateKey, err = os.ReadFile(utils.GetConfig().MustGet("oauth.private_key_path").MustString())
	if err != nil {
		panic(err)
	}

	publicKey, err = os.ReadFile(utils.GetConfig().MustGet("oauth.public_key_path").MustString())
	if err != nil {
		panic(err)
	}
}

func CreateAccessToken(user models.User) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		utils.GetLogger().Event("parse-key-error").Field("error", err).Warning()
		return "", CertKeyError.New("can not parse private key")
	}

	claims := AccessTokenClaims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "xybor.space",
			ExpiresAt: time.Now().Add(accessTokenExp).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		utils.GetLogger().Event("create-token-error").
			Field("claims", claims).Field("error", err).Warning()
		return "", TokenError.New("can not create token")
	}

	return token, nil
}

func ValidateUAT(token string) (AccessTokenClaims, error) {
	claims := AccessTokenClaims{}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		utils.GetLogger().Event("parse-key-error").Field("error", err).Warning()
		return claims, CertKeyError.New("can not parse public key")
	}

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, AlgorithmError.Newf("unexpected algorithm %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		if errors.Is(err, TokenError) {
			return claims, err
		}
		utils.GetLogger().Event("parse-token-error").Field("error", err).Debug()
		return claims, TokenError.New("can not parse the token")
	}

	return *parsedToken.Claims.(*AccessTokenClaims), nil
}
