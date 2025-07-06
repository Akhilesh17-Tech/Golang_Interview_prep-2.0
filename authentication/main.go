package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("akhil_secret_key_1705")

// Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to generate JWT
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there's an error while signing the token, return an empty string and the error.
		return "", fmt.Errorf("error generating JWT: %w", err)
	}

	// Return the signed token string.
	return tokenString, nil
}

// Middleware to authenticate JWT
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		if creds.Username == "akhil" && creds.Password == "akhilesh@07" {
			token, err := GenerateJWT(creds.Username)
			if err != nil {
				c.JSON(500, gin.H{"error": "could not generate token"})
				return
			}
			c.JSON(200, gin.H{"token": token})
			return
		}
		c.JSON(401, gin.H{"error": "invalid credentials"})
	})

	r.GET("/protected", AuthenticateJWT(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(200, gin.H{"message": "Hello " + username})
	})

	r.Run(":8080")
}
