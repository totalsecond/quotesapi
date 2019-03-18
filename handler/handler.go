package handler

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

// Quotes struct
type Quotes struct {
	RBCP []struct {
		ID    int    `json:"id"`
		Score int    `json:"score"`
		Text  string `json:"text"`
	} `json:"rbcp"`
}

// Quotez struct to variable
var Quotez Quotes

// GetQuote gets and returns all quotes in the db
func GetQuote(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Quotez.RBCP)
}

// GetRandom returns a random quote from the db
func GetRandom(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	n := rand.Int() % len(Quotez.RBCP)
	c.JSON(http.StatusOK, Quotez.RBCP[n])
}

// PostUpvote handles upvotes
func PostUpvote(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if quoteid, err := strconv.Atoi(c.Param("quoteID")); err == nil {
		for i := 0; i < len(Quotez.RBCP); i++ {
			if Quotez.RBCP[i].ID == quoteid {
				Quotez.RBCP[i].Score++
			}
		}
		theQuote := quoteid - 1
		c.JSON(http.StatusOK, &Quotez.RBCP[theQuote])
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// PostDownvote handles downvotes
func PostDownvote(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if quoteid, err := strconv.Atoi(c.Param("quoteID")); err == nil {
		for i := 0; i < len(Quotez.RBCP); i++ {
			if Quotez.RBCP[i].ID == quoteid {
				Quotez.RBCP[i].Score--
			}
		}
		theQuote := quoteid - 1
		c.JSON(http.StatusOK, &Quotez.RBCP[theQuote])
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
