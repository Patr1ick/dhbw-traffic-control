package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Patr1ick/dhbw-traffic-control/server/controller"
	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora/v3"
	"github.com/yugabyte/gocql"
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
	cluster := gocql.NewCluster(*settings.DatabaseAddress)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}
	cluster.Timeout = time.Second * 60
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(aurora.Red("Could not connect to Yugabyte on the first try..."))
		for i := 5; i > 0; i-- {
			log.Printf("Trying again (%v left)...\n", i)
			session, err := cluster.CreateSession()
			if err == nil {
				log.Printf("%s", aurora.Green("Connected to Database"))
				return session
			}
			time.Sleep(5 * time.Second)
		}
		log.Fatalln(aurora.Red("Could not connect to Database. Terminating server..."))
	}
	log.Printf("%s", aurora.Green("Connected to Yugabyte"))

	return session

}

func initDB(session *gocql.Session) {
	// Initialise KeySpace
	if err := session.Query("CREATE KEYSPACE IF NOT EXISTS traffic_control;").Exec(); err != nil {
		log.Fatal(aurora.Red("Could not create keyspace."))
	}

	// Initialise Table and Constraints
	var stmt = `
		CREATE TABLE IF NOT EXISTS traffic_control.clients ( x int, y int, z int, id uuid PRIMARY KEY ) WITH transactions = {'enabled': 'true'};
		CREATE UNIQUE INDEX IF NOT EXISTS traffic_control_coordinates ON traffic_control.clients(x, y, z);
	`

	if err := session.Query(stmt).Exec(); err != nil {
		log.Fatal(aurora.Red("Could not create database and constraint."))
	}
}

func Start(settings *model.Settings) {
	session := connectDB(settings)

	initDB(session)

	r := setupRouter(session, settings)
	r.Run()
}
