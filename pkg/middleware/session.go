package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Emmrys-Jay/altschool-sms/internal/config"
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Admin is the middleware for admin-only endpoints
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get and verify token
		token := c.GetHeader("Authorization")
		matNum, err := TokenIsVerified(token)
		if err != nil {
			rd := utility.BuildErrorResponse(http.StatusUnauthorized, "failed", "unauthorized", err.Error(), nil)
			c.JSON(http.StatusUnauthorized, rd)
			c.Abort()
			return
		}

		// Get information of the user sending request
		db := mysql.GetDB()
		if _, err := db.GetStudentByMatNum(c, matNum); err != nil {
			rd := utility.BuildErrorResponse(http.StatusUnauthorized, "failed", "unauthorized", err.Error(), nil)
			c.JSON(http.StatusUnauthorized, rd)
			c.Abort()
			return
		}

		// Check if request is not from an admin
		if matNum != config.GetConfig().AdminMatNum {
			rd := utility.BuildErrorResponse(http.StatusUnauthorized, "failed", "unauthorized", "cannot access admin endpoint", nil)
			c.JSON(http.StatusUnauthorized, rd)
			c.Abort()
			return
		}

		// Execute next handler
		c.Next()
	}
}

// Student is the middleware for all students endpoints
func Student() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get and verify token
		token := c.GetHeader("Authorization")
		matNum, err := TokenIsVerified(token)
		if err != nil {
			rd := utility.BuildErrorResponse(http.StatusUnauthorized, "failed", "unauthorized", err.Error(), nil)
			c.JSON(http.StatusUnauthorized, rd)
			c.Abort()
			return
		}

		// Get information of the user sending request
		db := mysql.GetDB()
		std, err := db.GetStudentByMatNum(c, matNum)
		if err != nil {
			rd := utility.BuildErrorResponse(http.StatusUnauthorized, "failed", "unauthorized", err.Error(), nil)
			c.JSON(http.StatusUnauthorized, rd)
			c.Abort()
			return
		}

		// Check if request is not from an admin
		if std.MatricNumber != config.GetConfig().AdminMatNum {
			if c.Param("id") != fmt.Sprint(std.ID) {
				rd := utility.BuildErrorResponse(http.StatusUnauthorized, "failed", "unauthorized", "access denied: unauthorized to access this resource", nil)
				c.JSON(http.StatusUnauthorized, rd)
				c.Abort()
				return
			}
		}

		// Execute next handler
		c.Next()
	}
}

func CreateToken(matNum string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti": matNum,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(config.GetConfig().SecretKey))
}

func TokenIsVerified(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().SecretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", errors.New("expired token")
		} else {
			return "", errors.New("invalid token")
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return "", errors.New("invalid token")
	}

	return claims["jti"].(string), nil
}
