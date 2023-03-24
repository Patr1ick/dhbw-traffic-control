package server

import (
	"log"
	"net/http"
	"os"

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

func connectDB() *gocql.Session {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "traffic_control"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(aurora.Red("Could not connect to Casssandra"))
		os.Exit(2)
	}
	log.Printf("%s on KeySpace %v", aurora.Green("Connected to Cassandra"), cluster.Keyspace)

	return session

}

func initDB(session *gocql.Session, settings *model.Settings) {
	// Check if area is existing
	var length int
	err := session.Query("SELECT COUNT(*) FROM traffic_area").Scan(&length)
	if err != nil {
		log.Fatalln(aurora.Red("Could not read from table"))
		os.Exit(3)
	}
	if length == 0 {
		log.Println(aurora.Yellow("No data in table. Initialise table..."))
		ta := &model.TrafficArea{
			Settings: settings,
		}
		ta.Create()
		// Save to DB
		X, Y, Z, Value := ta.ToTable()

		for i, val := range Value {
			err = session.Query(`INSERT INTO traffic_area (x, y, z, value) VALUES (?, ?, ?, ?)`, X[i], Y[i], Z[i], val).Exec()
			if err != nil {
				os.Exit(3)
			}
		}

	}
}

func Start(settings *model.Settings) {
	session := connectDB()

	initDB(session, settings)

	r := setupRouter(session, settings)
	r.Run()
}
