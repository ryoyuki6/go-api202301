package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "strconv"
)

type HealthData struct {
	ID		int		`json:"id"`
	Month	int		`json:"month"`
	Day		int		`json:"day"`
	Weight	float64	`json:"weight"`
}

var health_data = []HealthData{
	{ID:1, Month:1,	Day:1, Weight:55.5},
	{ID:2, Month:1,	Day:2, Weight:66.6},
	{ID:3, Month:1,	Day:3, Weight:77.7},
}

func getHealthdata(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, health_data)
}

func getHealthdataById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	for _, t := range health_data {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "the ID data does not exit!!!"})
}

func postHealthdata(c *gin.Context) {
	var newHealthdata HealthData
	if err := c.BindJSON(&newHealthdata); err != nil {
		return
	}
	health_data = append(health_data, newHealthdata)
}

func main() {
	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/health_data", getHealthdata)
	router.GET("/health_data/:id", getHealthdataById)
	router.POST("/health_data", postHealthdata)
	router.Run("localhost:8080")
}
