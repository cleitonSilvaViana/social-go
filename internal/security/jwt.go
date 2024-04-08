package security

import (
	"strings"

	"github.com/cleitonSilvaViana/social-go/config"
	"github.com/golang-jwt/jwt"
)

// UserClaims is used for set keys and values that contains datas about the user.
type UserClaims struct {
	Nick string `json:"nick"`
	jwt.StandardClaims
}

func extractSecretKey(token *jwt.Token) (interface{}, error) {
	return config.JWT_SECRET_KEY, nil
}

// NewAccessToken is used for create an JWT token.
func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString(config.JWT_SECRET_KEY)
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString(config.JWT_SECRET_KEY)
}

// ParsedAccessToken is used to extract header and body of token.
//
// Params:
//
//	accessToken is the access token for user.
//
// Return:
//
//	An strucutre of type UserClaims with the fields setted with datas extracted of access token
func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, extractSecretKey)
	return parsedAccessToken.Claims.(*UserClaims)
}

// ParseRefreshToken is used to ...
//
// Params:
//
// Return:
//
//	an structure with data validation of token
func ParseRefreshToken(refresToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refresToken, &jwt.StandardClaims{}, extractSecretKey)
	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}

// PreperToken is used to remove the substring "Bearer " of access token
//
// Params:
//
//	accessToken is an string with contains the access token of user.
//
// Return:
//
//	An access token withoud the substring "Bearer "
func PreperToken(accessToken string) string {
	if strings.Contains(accessToken, "Bearer") {
		splitToken := strings.Split(accessToken, "Bearer ")
		return splitToken[1]
	}
	return accessToken
}
