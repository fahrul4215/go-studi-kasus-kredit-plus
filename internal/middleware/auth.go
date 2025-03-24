package middleware

import (
	"fmt"
	config "go-studi-kasus-kredit-plus/configs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if open access
		if config.AppConfig.OpenAccess {
			c.Next()
			return
		}

		// Extract token from the Authorization header (Bearer <token>)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Token should start with "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Ensure token format is valid
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		// Parse the token using the secret key
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify signing method is HMAC with SHA256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Extract claims (user ID and roles) from the token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("roles", claims["roles"])
		}

		c.Next()
	}
}

// RoleMiddleware checks if a user has the required role.
func RoleMiddleware(requiredRole []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if open access
		if config.AppConfig.OpenAccess {
			c.Next()
			return
		}

		roles, exists := c.Get("roles")
		if !exists || roles == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient role"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		for _, requiredRole := range requiredRole {
			for _, role := range roles.([]interface{}) {
				if role.(string) == requiredRole {
					c.Next() // Allow request to proceed
					return
				}
			}
		}
		logrus.Warnf("User does not have the required role(s)")
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient role"})
		c.Abort()
	}
}
