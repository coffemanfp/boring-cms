package gin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/coffemanfp/test/config"
	dbErrors "github.com/coffemanfp/test/database/errors"
	sErrors "github.com/coffemanfp/test/server/errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// newCors creates a new CORS middleware using the provided configuration.
func newCors(conf config.ConfigInfo) gin.HandlerFunc {
	return cors.New(cors.Config{
		// Define allowed origins from the server configuration
		AllowOrigins: conf.Server.AllowedOrigins,
		// Define allowed HTTP methods
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		// Define allowed HTTP headers, including custom ones like "Authorization"
		AllowHeaders: []string{"*", "Authorization"},
		// Define headers exposed to clients in responses
		ExposeHeaders: []string{"Content-Length"},
		// Allow credentials (cookies, HTTP authentication) to be included in requests
		AllowCredentials: true,
		// Set the maximum amount of time that a preflight request can be cached
		MaxAge: 12 * time.Hour,
	})
}

// logger creates a Gin middleware for structured logging.
func logger() gin.HandlerFunc {
	// Use the structuredLogger function with a provided log.Logger instance
	return structuredLogger(&log.Logger)
}

// structuredLogger creates a Gin middleware that logs HTTP requests in a structured format.
func structuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time of the request handling
		start := time.Now()
		// Get the requested path and raw query parameters from the request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Continue processing the request
		c.Next()

		// Prepare parameters for log formatting
		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		// If raw query exists, append it to the path
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Choose log event type based on HTTP status code
		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		// Log the request details in a structured format
		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg(param.ErrorMessage)
	}
}

// authorize creates a Gin middleware that authorizes incoming requests based on a JWT token.
func authorize(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Attempt to save and validate token content
		err := saveTokenContent(c, secretKey)
		if err != nil {
			// If token content validation fails, return an unauthorized error
			err = sErrors.NewHTTPError(http.StatusUnauthorized, sErrors.UNAUTHORIZED_ERROR_MESSAGE)
			c.Error(err)
			c.Abort()
			return
		}

		// If token content validation succeeds, continue processing the request
		c.Next()
	}
}

// errorHandler creates a Gin middleware that handles errors and formats them into appropriate responses.
func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Iterate through the list of Gin errors
		for _, ginErr := range c.Errors {
			var isInternal bool

			// Check if the error is a custom database error
			if err, ok := ginErr.Err.(dbErrors.Error); ok {
				switch err.Type {
				case dbErrors.ALREADY_EXISTS:
					// If the error type is ALREADY_EXISTS, respond with a conflict status and message
					c.JSON(http.StatusConflict, gin.H{
						"message": sErrors.ALREADY_EXISTS,
					})

				case dbErrors.NOT_FOUND:
					// If the error type is NOT_FOUND, respond with a not found status and message
					c.JSON(http.StatusNotFound, gin.H{
						"message": sErrors.NOT_FOUND_ERROR_MESSAGE,
					})

				case dbErrors.UNKNOWN:
					isInternal = true
				}
			} else if err, ok := ginErr.Err.(sErrors.HTTPError); ok {
				// Check if the error is a custom HTTP error
				// Respond with the specified HTTP status code and error message
				c.JSON(err.Code, gin.H{
					"message": err.Message,
				})
			} else {
				// If the error is not recognized, consider it an internal server error
				isInternal = true
			}

			// Handle internal server errors or log other errors
			if isInternal {
				log.Error().Err(ginErr.Err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": sErrors.INTERNAL_SERVER_ERROR_MESSAGE,
				})
			} else {
				log.Info().Err(ginErr.Err)
			}
		}
	}
}

// readToken extracts the authentication token from various sources (query parameter or request header).
// It returns the extracted token or an error if no token is provided.
func readToken(c *gin.Context) (token string, err error) {
	// Check if the token is provided as a query parameter
	token = c.Query("token")
	if token != "" {
		return
	}

	// If the token is not found in the query parameter, check the "Authorization" header in the request
	token = c.Request.Header.Get("Authorization")
	if v := strings.Split(token, " "); len(v) == 2 {
		token = v[1]
	}

	// If no token is found in the query parameter or header, return an error
	if token == "" {
		err = errors.New("no token provided")
	}
	return
}

// saveTokenContent extracts and verifies the authentication token, and then extracts the account ID from the token's claims.
// It sets the extracted account ID in the Gin context.
func saveTokenContent(c *gin.Context, secretKey string) (err error) {
	// Read the authentication token from the request context
	tokenS, err := readToken(c)
	if err != nil {
		return
	}

	// Parse and verify the authentication token using the provided secret key
	token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}

	// Extract the claims from the token and ensure it's valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = errors.New("invalid token")
		return
	}

	// Extract the client ID from the claims and convert it to an integer
	id, ok := claims["client_id"].(float64)
	if !ok {
		err = errors.New("invalid token")
		return
	}

	// Set the extracted client ID in the Gin context for later use
	c.Set("id", int(id))
	return
}
