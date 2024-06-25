package test

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestDataloader(t *testing.T) {
	r := gin.Default()
	r.StaticFile("/data.csv", "./imdb_top_1000.csv")

	data := []map[string]interface{}{
		{"name": "Alice", "age": 18},
		{"name": "Bob", "age": 19},
		{"name": "Albert", "age": 20},
		{"name": "Clearance and Steve", "age": 18},
	}
	r.GET("/ndjson", func(c *gin.Context) {
		c.JSON(200, data)
	})
	r.Run("0.0.0.0:8080")
}
