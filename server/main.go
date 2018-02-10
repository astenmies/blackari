package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	mongo "./mongo"
	"github.com/graphql-go/handler"
	"github.com/nilstgmd/graphql-starter-kit/schema"
	"github.com/spf13/viper"

	_ "net/http/pprof"
)

func main() {
	// 1 Get the global config
	var (
		appName = viper.Get("app-name").(string)
		appHost = viper.Get("port").(string)
	)

	// Creates a GraphQL-go HTTP handler with the defined schema
	handler := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", handler)

	// serve a graphiql IDE
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	go func() {
		log.Println("Starting", appName, "on http://localhost:"+appHost)
		http.ListenAndServe(":9000", nil)
	}()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		for _ = range signalChan {
			log.Println("Received an interrupt, stopping GraphQL Server...")
			cleanup()
			cleanupDone <- true
		}
	}()

	<-cleanupDone
}

func init() {
	// Initialize viper
	// We can then call viper.Get("string") anywhere
	viper.SetConfigName("Config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	mongo.Init()
	// cassandra.Init()
}

func cleanup() {
	mongo.Cleanup()
	// cassandra.Cleanup()
}
