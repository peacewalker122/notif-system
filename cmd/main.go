package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"notifsys/internal/config"
	"notifsys/internal/factory"
	"notifsys/internal/server"
	"notifsys/pkg/db"
	"notifsys/pkg/fcm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	donechan := make(chan struct{})
	godotenv.Load(".env.local")
	cfg := config.Get().APP
	r := gin.New()
	database := db.New()

	err := fcm.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	f := factory.New(database.DB, donechan)

	server.Run(r, f)
	host := fmt.Sprintf(":%s", cfg.Host)

	go func() {
		err := http.ListenAndServe(host, r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Listening on %s", host)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)

	<-done
	close(donechan)
}
