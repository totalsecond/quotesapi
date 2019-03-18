package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"github.com/totalsecond/quotesapi/handler"
	//"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	rbcp := r.Group("/quotes/rbcp")
	{
		rbcp.GET("/", handler.GetQuote)
		rbcp.GET("/random", handler.GetRandom)
		rbcp.POST("/:quoteID/upvote", handler.PostUpvote)
		rbcp.POST("/:quoteID/downvote", handler.PostDownvote)
	}
	return r
}

func Start() {
	router := setupRouter()
	router.Run(":8080")
}
