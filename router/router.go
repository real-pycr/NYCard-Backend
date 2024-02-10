package router

import (
	"NYCard_Backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Error)

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/register", ctr.User.Register)
		apiRouter.POST("/login", ctr.User.Login)
		apiRouter.DELETE("/logout", ctr.User.Logout)
		apiRouter.GET("/getinfo", ctr.User.GetUserStatus)

		apiRouter.POST("/addcard", ctr.Card.AddCard)
		apiRouter.GET("/getcard", ctr.Card.GetCard)
		apiRouter.DELETE("/deletecard/:id", ctr.Card.DeleteCard)
		apiRouter.GET("/getcardlist", ctr.Card.GetCardList)

		apiRouter.POST("/addpoem", ctr.Poem.AddPoem)
		apiRouter.GET("/getpoem", ctr.Poem.GetPoem)
	}

}
