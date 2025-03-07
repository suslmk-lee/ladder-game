package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type LadderRequest struct {
	Participants []string `json:"participants"`
	Outcomes     []string `json:"outcomes"`
}

type LadderStructure struct {
	Ladder [][]int `json:"ladder"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	// 헬스 체크 엔드포인트 추가
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/generate", func(c *gin.Context) {
		var req LadderRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// 간단한 사다리 구조 생성 (참가자 수에 따라)
		numParticipants := len(req.Participants)
		ladder := make([][]int, numParticipants-1) // 가로선은 세로선 사이에 위치
		for i := range ladder {
			ladder[i] = make([]int, numParticipants)
			// 랜덤으로 가로선 추가 (0 또는 1)
			for j := 0; j < numParticipants-1; j++ {
				ladder[i][j] = rand.Intn(2) // 50% 확률로 가로선 생성
			}
		}

		c.JSON(http.StatusOK, LadderStructure{Ladder: ladder})
	})

	r.Run(":8080")
}
