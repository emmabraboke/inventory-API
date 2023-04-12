package middlewares

import (
	"inventory/internals/service/tokenService"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authMiddlewareSrv struct {
	tokenSrv tokenService.TokenService
}

func NewAuthMiddleware(tokenSrv tokenService.TokenService) *authMiddlewareSrv {
	return &authMiddlewareSrv{tokenSrv: tokenSrv}

}
func (t *authMiddlewareSrv) AuthMiddleware(c *gin.Context) {

	// Check if the "Authorization" header is present
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "unathorized")
		return
	}

	// Verify the authentication token
	auth := strings.Split(authHeader, " ")

	if auth[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
		return
	}

	token := auth[1] // Remove the "Bearer " prefix from the header value

	// var claims t.
	claims, err := t.tokenSrv.ValidateToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", claims.Id)
	c.Set("email", claims.Email)

	c.Next()

}
