package middlewares

import (
	"log"
	"net/http"

	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")

		if len(authHeader) == 0 {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		token := authHeader[len(BEARER_SCHEMA):]

		auth, err := services.NewJWTService().ValidateToken(token)

		if auth.Valid {
			claims := auth.Claims.(jwt.MapClaims)
			log.Println("Claims[Email]: ", claims["email"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
			ctx.Set("Email", claims["email"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
