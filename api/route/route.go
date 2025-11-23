package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quandev306/webcrawl/bootstrap"
	"github.com/quandev306/webcrawl/config"
)

// Setup registers all routes. Add more route groups here as you build features.
func Setup(env *bootstrap.Env, timeout time.Duration, db config.Database, r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"env":    env.AppEnv,
		})
	})
	_ = timeout
	_ = db
}
