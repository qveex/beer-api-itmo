package jwt

import (
	"fmt"
	"time"

	"grpc-service-ref/internal/domain/models"

	"github.com/golang-jwt/jwt/v5"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ExtractUserInfoFromJWT извлекает email и uid из JWT токена.
// Возвращает UserDto и ошибку, если токен неверен или не содержит email/uid.
func ExtractUserInfoFromJWT(tokenString string, secretKey string) (string, error) {
	// Парсинг и верификация токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %w", err)
	}

	// Проверка валидности токена и извлечение claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["email"].(string)
		//uid, ok := claims["uid"].(int64)
		if !ok {
			return "", fmt.Errorf("email or uid not found in token")
		}
		return email, nil
	}

	return "", fmt.Errorf("invalid JWT token")
}
