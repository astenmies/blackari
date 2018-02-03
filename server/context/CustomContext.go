package context

import (
	"context"
	"log"
	"time"

	h "github.com/OscarYuen/go-graphql-starter/handler"
	"github.com/OscarYuen/go-graphql-starter/service"
	"github.com/spf13/viper"
)

// CustomContext uses context, which makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.
// The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request's deadline.
// ref https://blog.golang.org/context
// ref https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39
func CustomContext() context.Context {

	viper.SetConfigName("Config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	var (
		appName             = viper.Get("app-name").(string)
		signedSecret        = viper.Get("auth.jwt-secret").(string)
		expiredTimeInSecond = time.Duration(viper.Get("auth.jwt-expire-in").(int64))
		debugMode           = viper.Get("log.debug-mode").(bool)
		logFormat           = viper.Get("log.log-format").(string)
	)

	ctx := context.Background()
	log := h.NewLogger(&appName, debugMode, &logFormat)
	// roleService := service.NewRoleService(db, log)
	// userService := service.NewUserService(db, roleService, log)
	authService := service.NewAuthService(&appName, &signedSecret, &expiredTimeInSecond, log)

	ctx = context.WithValue(ctx, "log", log)
	// ctx = context.WithValue(ctx, "roleService", roleService)
	// ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "authService", authService)

	return ctx
}
