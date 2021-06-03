package main

import (
	"accountServer/config"
	"accountServer/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "accountServer config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// set gin mode
	gin.SetMode(viper.GetString("runmode"))

	// create the gin engine
	g := gin.New()

	// gin middlewares
	middlewares := []gin.HandlerFunc{}

	// Routes
	router.Load(
		g,              // cores
		middlewares..., // middlewares
	)

	// ping the svr to make sure the router is working
	go func() {
		if err := ping(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
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
