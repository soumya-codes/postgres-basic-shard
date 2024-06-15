package router

import (
	"github.com/gin-gonic/gin"
	"github.com/soumya-codes/postgres-static-shard/internal/handler/heartbeat"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/heartbeat/:machine_id", heartbeat.GetLastHeartbeat)
	r.PUT("/heartbeat/:machine_id", heartbeat.UpdateLastHeartbeat)
	return r
}
