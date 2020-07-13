package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	_ "Gogin/internal/helpers/environment"
	circuitBreaker "Gogin/internal/platform/circuitBreaker"
	_ "Gogin/internal/platform/database"
	_ "Gogin/internal/platform/redis"

	authentication "Gogin/api/Authentication"
	identity "Gogin/api/Identity"

	"github.com/kylegrantlucas/speedtest"
)

var g errgroup.Group

func routes() http.Handler {
	r := gin.New()

	// Hystrix
	r.Use(circuitBreaker.HystrixHandler("gogin"))

	// Recover
	r.Use(gin.Recovery())

	// Load static resources
	r.LoadHTMLGlob("./views/*.html")
	r.Static("/css", "./views/css")
	r.Static("/js", "./views/js")
	r.Static("/img", "./views/img")

	// Homepage
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// ping
	r.GET("/ping", func(c *gin.Context) {
		client, err := speedtest.NewDefaultClient()
		if err != nil {
			fmt.Printf("error creating client: %v", err)
		}

		// Pass an empty string to select the fastest server
		server, err := client.GetServer("")
		if err != nil {
			fmt.Printf("error getting server: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"country":  server.Country,
			"server":   server.Name,
			"distance": fmt.Sprintf("%3.2f km", server.Distance),
			"latency":  fmt.Sprintf("%3.2f ms", server.Latency),
		})

	})

	// Routes
	authentication.Routes(r)
	identity.Routes(r)

	return r
}

func main() {
	// Multiple instances for load balancing
	for _, i := range []int{0, 1, 2} {
		server := &http.Server{
			Addr:         ":" + os.Getenv("PORT"+strconv.Itoa(i)),
			Handler:      routes(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		g.Go(func() error {
			return server.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
