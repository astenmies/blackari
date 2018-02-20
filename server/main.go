package main

import (
	mongo "./mongo"
	"./resolver"
	"./schema"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

var graphqlSchema *graphql.Schema

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

	// Creates a GraphQL-go HTTP handler with the defined schema
	graphqlSchema = graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

}

func cleanup() {
	mongo.Cleanup()
}

func main() {
	// 1 Get the global config
	var (
		appName = viper.Get("app-name").(string)
		appHost = viper.Get("port").(string)
	)

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", cors.Default().Handler(&relay.Handler{Schema: graphqlSchema}))
	http.Handle("/query", &relay.Handler{Schema: graphqlSchema})

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
