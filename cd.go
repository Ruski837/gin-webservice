package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cds = []album{
	{ID: "1", Title: "Take off your pants and jacket", Artist: "Blink 182", Price: 9.99},
	{ID: "2", Title: "Metoria", Artist: "Linken Park", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getCds(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cds)
}
