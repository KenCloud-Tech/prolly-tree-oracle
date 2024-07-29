package test

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestDataloader(t *testing.T) {
	r := gin.Default()
	r.StaticFile("/data.csv", "./imdb_top_1000.csv")
	r.StaticFile("/ndjson", "./ndjson")

	r.Run("0.0.0.0:8080")
}
