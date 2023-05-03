package api

import (
	db "Gym-backend/db/sqlc"
	"Gym-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createCheckinActivityRequest struct {
	Checkin    time.Time `json:"checkin" binding:"required"`
	Checkout   time.Time `json:"checkout" binding:"required"`
	Userid     int64     `json:"userid" binding:"required"`
	Employeeid int64     `json:"employeeid" binding:"required"`
	Locationid int64     `json:"locationid" binding:"required"`
}

type createCheckinRecordRequest struct {
	Type       int32  `json:"type" binding:"required"`
	Useremail  string `json:"useremail" binding:"required"`
	Employeeid int64  `json:"employeeid" binding:"required"`
	Locationid int64  `json:"locationid" binding:"required"`
}

type getCheckinActivityRequest struct {
	Userid int64 `form:"userid" binding:"required"`
}

type checkinActivityResponse struct {
	ID         int64     `json:"id"`
	Checkin    time.Time `json:"checkin"`
	Checkout   time.Time `json:"checkout"`
	Userid     int64     `json:"userid"`
	Employeeid int64     `json:"employeeid"`
	Locationid int64     `json:"locationid"`
}

func newCheckinActivityResponse(activity db.Checkinactivity) checkinActivityResponse {
	return checkinActivityResponse{
		ID:         activity.ID,
		Checkin:    activity.Checkin,
		Checkout:   activity.Checkout,
		Userid:     activity.Userid,
		Employeeid: activity.Employeeid,
		Locationid: activity.Locationid,
	}
}

// CreateTags		godoc
// @Summary			Create CheckinActivity
// @Description 	Create CheckinActivity data in Db.
// @Param 			device body createCheckinRecordRequest true "Create Checkin Activity Record"
// @Produce 		application/json
// @Tags 			checkinActivity
// @Success 		200 {object} string
// @Router			/checkinRecord [post]
func (server *Server) createCheckinRecord(ctx *gin.Context) {

	var req createCheckinRecordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserFromEmail(ctx, req.Useremail)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateCheckinRecordsParams{
		Userid:     user.ID,
		Type:       req.Type,
		Employeeid: req.Employeeid,
		Locationid: req.Locationid,
	}

	_, err = service.CreateCheckInRecord(ctx, server.store, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Checkin Recorder")
}

// CreateTags		godoc
// @Summary			Create CheckinActivity
// @Description 	Create CheckinActivity data in Db.
// @Param 			device body createCheckinRecordRequest true "Create Checkin Activity Record"
// @Produce 		application/json
// @Tags 			checkinActivity
// @Success 		200 {object} checkinActivityResponse{}
// @Router			/checkoutRecord [post]
func (server *Server) createCheckOutRecord(ctx *gin.Context) {
	var req createCheckinRecordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserFromEmail(ctx, req.Useremail)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateCheckinRecordsParams{
		Userid:     user.ID,
		Type:       req.Type,
		Employeeid: req.Employeeid,
		Locationid: req.Locationid,
	}

	activity, err := service.CreateCheckOutRecord(ctx, server.store, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newCheckinActivityResponse(activity)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Create CheckinActivity
// @Description 	Create CheckinActivity data in Db.
// @Param 			device body createCheckinActivityRequest true "Create Checkin Activity"
// @Produce 		application/json
// @Tags 			checkinActivity
// @Success 		200 {object} checkinActivityResponse{}
// @Router			/checkinActivity [post]
func (server *Server) createCheckinActivity(ctx *gin.Context) {
	var req createCheckinActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateCheckinActivityParams{
		Checkin:    req.Checkin,
		Checkout:   req.Checkout,
		Userid:     req.Userid,
		Employeeid: req.Employeeid,
		Locationid: req.Locationid,
	}

	activity, err := server.store.CreateCheckinActivity(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newCheckinActivityResponse(activity)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get CheckinActivity
// @Description 	Get CheckinActivity from  data in Db.
// @Param 			activity query getCheckinActivityRequest true "Get Checkin Activity"
// @Produce 		application/json
// @Tags 			checkinActivity
// @Success 		200 {object} []checkinActivityResponse{}
// @Router			/checkinActivity [get]
func (server *Server) getCheckinActivity(ctx *gin.Context) {
	var req getCheckinActivityRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	activity, err := server.store.GetCheckinActivity(ctx, req.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, activity)
}
