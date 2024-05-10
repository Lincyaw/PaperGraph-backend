package handlers

import (
	"context"
	_ "github.com/Lincyaw/PaperGraph-backend/config"
	"github.com/Lincyaw/PaperGraph-backend/drivers"
	"github.com/Lincyaw/PaperGraph-backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"net/http"
)

// Paper 函数处理 "/paper?query=xxx" 路径的 GET 请求
func Paper(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, NewErrResponse(ErrBadRequest, "query parameter need", nil))
		return
	}
	cypherQuery := `MATCH (n) WHERE toLower(n.title) CONTAINS toLower($query) RETURN n`
	logger.Debug(query)
	engine := drivers.GetInstance()
	result, err := engine.ExecuteRead(context.Background(), cypherQuery, map[string]interface{}{"query": query})
	if err != nil {
		logger.Error("query graph database", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResponse(err, "", nil))
		return
	}
	var results []string
	for _, v := range result.Records {
		m := v.AsMap()
		results = append(results, m["n"].(dbtype.Node).Props["title"].(string))
	}
	c.JSON(http.StatusOK, NewResponse(results))
}
