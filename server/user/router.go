package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Register(gr *gin.Engine)
}
type router struct {
	handler UserHandler
}

func (r *router) login(c *gin.Context) {
	var request BaseUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "error",
		})
		return
	}
	result, err := r.handler.Login(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, &BaseUserResponse{Token: result})
}

func (r *router) register(c *gin.Context) {
	var request BaseUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, map[string]string{
			"error_message": "error",
		})
		return
	}
	result, err := r.handler.Register(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, &BaseUserResponse{Token: result})
}

func (r *router) Register(gr *gin.Engine) {
	user := gr.Group("/user")
	{
		user.POST("/login", r.login)
		user.POST("/register", r.register)
	}
	internal := gr.Group("/internal")
	{
		internal.POST("/Verify", r.login)
	}
}
func NewRouter(handler UserHandler) UserRouter {
	return &router{
		handler: handler,
	}
}
