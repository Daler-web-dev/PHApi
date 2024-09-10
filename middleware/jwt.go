package middleware

import (
	"net/http"
	"simple-backend/utils"
	"strings"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Ожидаем, что токен будет в формате "Bearer <token>"
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Проверяем валидность токена
		_, err := utils.ParseToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// // Выводим username из токена для отладки (по желанию)
		// username := claims["username"]
		// fmt.Println("Authenticated user:", username)

		// Передаем управление следующему хендлеру
		next.ServeHTTP(w, r)
	})
}
