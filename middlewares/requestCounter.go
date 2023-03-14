package middlewares

import (
	"example/go-api/initializers"
	"example/go-api/models"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiCallCounter struct {
	mu    sync.Mutex
	Count int
}

func (cc *ApiCallCounter) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc.Count++
		c.Next()
	}
}

func (cc *ApiCallCounter) ResetCounter() {
	cc.mu.Lock()
	cc.Count = 0
	cc.mu.Unlock()
}

func (cc *ApiCallCounter) Start() {
	for {
		var startTime = time.Now().UTC()
		time.Sleep(1 * time.Hour)
		cc.mu.Lock()
		if cc.Count != 0 {
			initializers.DB.Create(&models.ApiCall{ApiCallsCount: cc.Count, StartTime: startTime, EndTime: time.Now().UTC()})
		}
		cc.mu.Unlock()
		cc.ResetCounter()
	}
}
