package main

import (
	"examples/identity/config"
	"examples/identity/db"
	"examples/identity/jwthelper"

	modeluser "examples/identity/model/user"
	internal "examples/identity/server/internalapi"
	serveruser "examples/identity/server/user"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	s := config.LoadEnvConfig()
	db := db.InitDatabase(s)

	userDatabase := modeluser.NewDatabase(db)
	userRepo := modeluser.NewRepository(userDatabase)
	jwtHandler := jwthelper.NewJWTHelper(s)
	userHandler := serveruser.NewHandler(userRepo, jwtHandler)
	userRouter := serveruser.NewRouter(userHandler)
	userRouter.Register(r)

	internalRouter := internal.NewInternalRouter(jwtHandler)
	internalRouter.Register(r)

	r.Run(":" + s.Port)
}
