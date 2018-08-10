package sd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message+"\n")
}
