package controller

import (
	"fmt"
	"net/http"

	"github.com/Patr1ick/dhbw-traffic-control/server/logic"
	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yugabyte/gocql"
)

func HandleClientStart(ctx *gin.Context, session *gocql.Session, settings *model.Settings) {

	payload := &model.PayloadCoordinates{}

	err := ctx.ShouldBind(&payload)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "could not read payload",
			},
		)
		return
	}

	cl, err := logic.LoadTable(session, settings)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "could not load table",
			},
		)
		return
	}

	id := uuid.New()
	cord := model.Coordinate{X: payload.Pos.X, Y: payload.Pos.Y}

	if !settings.Valid(cord) {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Errorf("no valid position"),
			},
		)
		return
	}

	client := &model.Client{
		Id:  id,
		Pos: cord,
	}

	if err = logic.InitClient(session, cl, client); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		client,
	)
}

func HandleClientMove(ctx *gin.Context, session *gocql.Session, settings *model.Settings) {
	payload := &model.PayloadMove{}

	err := ctx.ShouldBind(&payload)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "could not read payload",
			},
		)
		return
	}

	ta, err := logic.LoadTable(session, settings)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "could not load table",
			},
		)
		return
	}

	oldPos, newPos, err := logic.Move(ta, payload.Id, payload.Target)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "could not move",
			},
		)
		return
	}

	if err = logic.UpdateClient(model.Client{Pos: *newPos, Id: payload.Id}, session); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"internal": err.Error(),
				"message":  "failed to write to db",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"oldPos": oldPos,
			"newPos": newPos,
		},
	)
}
