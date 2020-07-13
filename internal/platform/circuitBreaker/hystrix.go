package circuitbreaker

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func init() {
	hystrix.ConfigureCommand("gogin", hystrix.CommandConfig{
		// How long to wait for command to complete, in milliseconds
		Timeout: 50000,

		// MaxConcurrent is how many commands of the same type
		// can run at the same time
		MaxConcurrentRequests: 300,

		// VolumeThreshold is the minimum number of requests
		// needed before a circuit can be tripped due to health
		RequestVolumeThreshold: 10,

		// SleepWindow is how long, in milliseconds,
		// to wait after a circuit opens before testing for recovery
		SleepWindow: 1000,

		// ErrorPercentThreshold causes circuits to open once
		// the rolling measure of errors exceeds this percent of requests
		ErrorPercentThreshold: 50,
	})
}

// HystrixHandler func
func HystrixHandler(command string) gin.HandlerFunc {
	return func(c *gin.Context) {
		hystrix.Do(command, func() error {
			c.Next()
			return nil
		}, func(err error) error {
			//Handle failures
			return err
		})
	}
}
