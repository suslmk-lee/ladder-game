package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LadderRequest struct {
	Participants []string `json:"participants"`
	Outcomes     []string `json:"outcomes"`
}

type LadderStructure struct {
	Ladder [][]int `json:"ladder"`
}

func main() {
	r := gin.Default()

	// 헬스 체크 엔드포인트 추가
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/run", func(c *gin.Context) {
		var req LadderRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// 백엔드 2(Ladder Generator)에 사다리 구조 요청
		jsonData, _ := json.Marshal(req)
		resp, err := http.Post("http://ladder-generator-service:8080/generate", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate ladder"})
			return
		}
		defer resp.Body.Close()

		var ladder LadderStructure
		json.NewDecoder(resp.Body).Decode(&ladder)

		// 백엔드 3(Result Mapper)에 결과 매핑 요청
		mappingReq := struct {
			LadderStructure
			LadderRequest
		}{ladder, req}
		jsonData, _ = json.Marshal(mappingReq)
		resp, err = http.Post("http://result-mapper-service:8080/map", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map results"})
			return
		}
		defer resp.Body.Close()

		var result map[string]string
		json.NewDecoder(resp.Body).Decode(&result)
		c.JSON(http.StatusOK, gin.H{"results": result})
	})

	r.Run(":8080")
}
