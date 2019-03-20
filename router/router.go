package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"github.com/totalsecond/quotesapi/handler"
	//"log"
)

// SetupRouter configures our router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(favicon.New("./views/favicon.ico"))
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
