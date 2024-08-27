package Infrastructure

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    Email string `json:"email"`
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

type TokenService struct {}

func NewTokenService() *TokenService {
    return &TokenService{}
}

func (s *TokenService) GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func (s *TokenService) ValidateToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, err
        }
        return nil, err
    }
    if !token.Valid {
        return nil, err
    }
    return claims, nil
}

//generates both an access token and a refresh token.
func (s *TokenService) GenerateTokens(userID string) (string, string, error) {
    accessToken, err := s.GenerateToken(userID)
    if err != nil {
        return "", "", err
    }

    refreshToken, err := s.GenerateToken(userID) 
    if err != nil {
        return "", "", err
    }

    return accessToken, refreshToken, nil
}