package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"notifsys/internal/config"
	"notifsys/internal/factory"
	"notifsys/internal/middleware"
	"notifsys/internal/server"
	"notifsys/pkg/db"
	"notifsys/pkg/fcm"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/uptrace/uptrace-go/uptrace"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	ctx := context.Background()

	donechan := make(chan struct{})
	godotenv.Load(".env.local")
	cfg := config.Get().APP
	r := gin.New()
	middleware.Run(r)
	database := db.New()

	err := fcm.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	f := factory.New(database.DB, donechan)

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(cfg.Name),
	)
	defer uptrace.Shutdown(ctx)

	server.Run(r, f)
	host := fmt.Sprintf(":%s", cfg.Host)

	if os.Getenv("MODE") == "DEBUG" {
		pprof.Register(r)
	}

	srv := &http.Server{
		Addr:    host,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Listening on %s", host)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)

	<-done
	close(donechan)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	database.Close()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout")
	}

	log.Println("shutting down")
}
