package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Patr1ick/dhbw-traffic-control/server/controller"
	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/logrusorgru/aurora/v3"
)

func setupRouter(session *gocql.Session, settings *model.Settings) *gin.Engine {
	r := gin.Default()
	// Routes
	v1 := r.Group("/v1")
	{
		traffic := v1.Group("/traffic")
		{
			traffic.POST("/start", func(ctx *gin.Context) {
				controller.HandleClientStart(ctx, session, settings)
			})
			traffic.POST("/move", func(ctx *gin.Context) {
				controller.HandleClientMove(ctx, session, settings)
			})
		}
		v1.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
	}

	return r
}

func connectDB(settings *model.Settings) *gocql.Session {
	cluster := gocql.NewCluster(*settings.CassandraAddress)
	cluster.Keyspace = "traffic_control"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(aurora.Red("Could not connect to Casssandra on the first try..."))
		for i := 5; i > 0; i-- {
			log.Printf("Trying again (%v left)...\n", i)
			session, err := cluster.CreateSession()
			if err == nil {
				log.Printf("%s on KeySpace %v", aurora.Green("Connected to Cassandra"), cluster.Keyspace)
				return session
			}
			time.Sleep(5 * time.Second)
		}
	}
	log.Printf("%s on KeySpace %v", aurora.Green("Connected to Cassandra"), cluster.Keyspace)

	return session

}

func Start(settings *model.Settings) {
	session := connectDB(settings)

	r := setupRouter(session, settings)
	r.Run()
}
