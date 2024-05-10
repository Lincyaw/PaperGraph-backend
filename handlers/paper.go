package handlers

import (
	"context"
	_ "github.com/Lincyaw/PaperGraph-backend/config"
	"github.com/Lincyaw/PaperGraph-backend/drivers"
	"github.com/Lincyaw/PaperGraph-backend/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Paper 函数处理 "/paper?query=xxx" 路径的 GET 请求
func Paper(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter is required"})
		return
	}
	cypherQuery := `MATCH (n) WHERE n.content CONTAINS $query RETURN n`
	logger.Debug(query)
	engine := drivers.GetInstance()
	result, err := engine.ExecuteRead(context.Background(), cypherQuery, map[string]interface{}{"query": query})
	if err != nil {
		logger.Error("query graph database", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResponse(err, "", nil))
		return
	}

	var nodes []map[string]interface{}

	for result.Next(context.Background()) {
		resultMap := result.Record().AsMap()
		nodes = append(nodes, resultMap)
	}

	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nodes})
}
