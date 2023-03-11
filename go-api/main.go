package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
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

func getHalthdata(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, health_data)
}

func postHalthdata(c *gin.Context) {
	var newHalthdata HealthData
	if err := c.BindJSON(&newHalthdata); err != nil {
		return
	}
	health_data = append(health_data, newHalthdata)
}

func main() {
	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/health_data", getHalthdata)
	router.POST("/health_data", postHalthdata)
	router.Run("localhost:8080")
}
