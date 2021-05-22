package main

import (
	"accountServer/router"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	g := gin.New()                    // Create the Gin engine
	middlewares := []gin.HandlerFunc{} // gin middlewares
	// Routes
	router.Load(
		g,              // cores
		middlewares..., // middlewares
	)

	// 自检
	go func() {
		if err := ping(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listen the coming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

// 自检保证启动后的API服务器处于健康状态
func ping() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sc/health")
		// svr正常运作退出ping()
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
