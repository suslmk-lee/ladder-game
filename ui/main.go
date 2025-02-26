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

func main() {
	r := gin.Default()

	// API 엔드포인트
	r.POST("/ladder", func(c *gin.Context) {
		var req LadderRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// 백엔드 1(Ladder Manager)에 요청
		jsonData, _ := json.Marshal(req)
		resp, err := http.Post("http://ladder-manager:8080/run", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contact backend"})
			return
		}
		defer resp.Body.Close()

		// 결과 처리
		var result map[string]string
		json.NewDecoder(resp.Body).Decode(&result)
		c.JSON(http.StatusOK, gin.H{"results": result})
	})

	// 정적 파일 제공 (React 빌드 파일)
	r.StaticFS("/static", http.Dir("./frontend/build/static"))
	r.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
	r.StaticFile("/manifest.json", "./frontend/build/manifest.json")
	
	// 모든 다른 경로는 index.html로 라우팅 (React Router 지원)
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	r.Run(":8080")
}
