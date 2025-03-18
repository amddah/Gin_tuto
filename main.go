package main

import (
	"gin_api/adapter/controllers"
	"gin_api/adapter/handlers"
	"gin_api/adapter/routes"
	"gin_api/core/events"
	"gin_api/initializer"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVaraibles()
	initializer.ConnectionToDB()
	initializer.ConnectToRedis()
	initializer.ConnectToMongoDB()
}

func main() {
	r := gin.Default()
	defer initializer.DisconnectMongoDB()

	// Initialize controllers after Redis connection is established
	controllers.InitUserController()
	controllers.InitPostController()

	// Set up event subscriber and handler for users
	userSubscriber := events.NewRedisEventSubscriber(initializer.RedisClient, "user-events")
	userEventHandler := handlers.NewUserEventHandler()
	userSubscriber.Subscribe(userEventHandler)
	defer userSubscriber.Close()

	// Set up event subscriber and handler for posts
	postSubscriber := events.NewRedisEventSubscriber(initializer.RedisClient, "post-events")
	postEventHandler := handlers.NewPostEventHandler()
	postSubscriber.Subscribe(postEventHandler)
	defer postSubscriber.Close()

	// Serve pprof in a separate goroutine
	go func() {
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()

	// pprof routes
	r.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	r.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	r.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	r.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))

	api := r.Group("/api")
	{
		routes.PostRoutes(api)
		routes.UserRoutes(api)
	}
	// Run Gin
	r.Run()
}
