package app

import (
	"encoding/gob"

	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *api.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/template", "../template/static")
	router.LoadHTMLGlob("../template/web/*")

	router.GET("/", HomeHandler)
	router.GET("/login", LoginHandler(auth))
	router.GET("/callback", CallBackHandler(auth))
	router.GET("/user", IsAuthenticated, UserHandler)
	router.GET("/logout", LogOutHandler)

	return router
}
