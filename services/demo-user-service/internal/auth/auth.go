package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

func GenerateJWT(userID string) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerBytes)

	claims := map[string]interface{}{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsBytes)

	unsignedToken := headerEncoded + "." + claimsEncoded
	mac := hmac.New(sha256.New, []byte(os.Getenv("JWT_SECRET")))
	mac.Write([]byte(unsignedToken))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return unsignedToken + "." + signature, nil
}

type contextKey string

const userIDContextKey contextKey = "userID"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDContextKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateJWT(tokenStr string) (string, error) {
	parts := strings.Split(tokenStr, ".")
	if len(parts) != 3 {
		return "", &jwtError{"invalid token format"}
	}

	unsignedToken := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, []byte(os.Getenv("JWT_SECRET")))
	mac.Write([]byte(unsignedToken))
	expectedSignature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return "", &jwtError{"invalid signature"}
	}

	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", &jwtError{"invalid claims encoding"}
	}

	var claims map[string]interface{}
	err = json.Unmarshal(claimsBytes, &claims)
	if err != nil {
		return "", &jwtError{"invalid claims format"}
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return "", &jwtError{"token expired"}
		}
	} else {
		return "", &jwtError{"invalid expiration"}
	}

	if sub, ok := claims["sub"].(string); ok {
		return sub, nil
	}

	return "", &jwtError{"invalid subject"}
}

type jwtError struct {
	s string
}

func (e *jwtError) Error() string {
	return e.s
}
