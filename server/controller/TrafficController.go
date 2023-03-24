package controller

import (
	"math/rand"
	"net/http"

	"github.com/Patr1ick/dhbw-traffic-control/server/logic"
	"github.com/Patr1ick/dhbw-traffic-control/server/model"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func HandleClientStart(ctx *gin.Context, session *gocql.Session, settings *model.Settings) {

	payload := &model.PayloadCoordinates{}

	err := ctx.ShouldBind(&payload)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
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

	id := rand.Intn(ta.Width*ta.Height*ta.Depth - 1)
	cord := model.Coordinate{X: payload.Pos.X, Y: payload.Pos.Y}

	pos, err := ta.Set(id, cord)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "could not set the id at the position",
				"error":   err.Error(),
			},
		)
		return
	}

	err = logic.SavePos(id, *pos, session)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "could not save the position",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"id":     id,
			"pos":    cord,
			"y_area": ta.Area,
		},
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

	if err = logic.SavePos(payload.Id, *newPos, session); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "failed to write to db",
				"oldPos":  oldPos,
				"newPos":  newPos,
				"err":     err.Error(),
			},
		)
		return
	}
	if err = logic.SavePos(-1, *oldPos, session); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "failed to write to db",
				"oldPos":  oldPos,
				"newPos":  newPos,
				"err":     err.Error(),
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
