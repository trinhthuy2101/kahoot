package internalapi

import (
	"examples/identity/jwthelper"

	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalRouter interface {
	Register(g *gin.Engine)
}
type router struct {
	jwtHelper jwthelper.JWTHelper
}

func (r *router) verifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request IdentityRequest
		c.ShouldBindJSON(&request)

		token, _ := r.jwtHelper.ValidateJWT(request.TokenString)
		if token.Valid {
			c.JSON(http.StatusOK, IdentityResponse{
				IsValid: true,
			})
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func (r *router) Register(g *gin.Engine) {
	internal := g.Group("/internal")
	{
		internal.POST("/VerifyToken", r.verifyToken())
	}
}
func NewInternalRouter(h jwthelper.JWTHelper) InternalRouter {
	return &router{
		jwtHelper: h,
	}
}
