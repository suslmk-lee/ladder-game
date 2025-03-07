package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LadderRequest struct {
	Participants []string `json:"participants"`
	Outcomes     []string `json:"outcomes"`
	Ladder       [][]int  `json:"ladder"`
}

func main() {
	r := gin.Default()

	// 헬스 체크 엔드포인트 추가
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/map", func(c *gin.Context) {
		var req LadderRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// 간단한 매핑 로직: 사다리 경로를 따라 결과 연결
		numParticipants := len(req.Participants)
		results := make(map[string]string)
		for i := 0; i < numParticipants; i++ {
			current := i
			for _, row := range req.Ladder {
				if current < numParticipants-1 && row[current] == 1 {
					current++ // 오른쪽으로 이동
				} else if current > 0 && row[current-1] == 1 {
					current-- // 왼쪽으로 이동
				}
			}
			results[req.Participants[i]] = req.Outcomes[current]
		}

		c.JSON(http.StatusOK, results)
	})

	r.Run(":8080")
}
