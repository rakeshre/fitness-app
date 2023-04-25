package api

import (
	db "Gym-backend/db/sqlc"
	"Gym-backend/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createUserActivityRequest struct {
	Start    string `json:"start" binding:"required"`
	End      string `json:"end" binding:"required"`
	Userid   int64  `json:"userid" binding:"required"`
	Deviceid int64  `json:"deviceid" binding:"required"`
}

type getUserActivityRequest struct {
	Userid int64 `form:"userid" binding:"required"`
}

type getPastWorkOutActivityRequest struct {
	Userid   int64  `form:"userid" binding:"required"`
	Interval string `form:"interval" binding:"required"`
}

type createUserActivityRecordRequest struct {
	Type       int32 `json:"type" binding:"required"`
	Userid     int64 `json:"userid" binding:"required"`
	Locationid int64 `json:"locationid" binding:"required"`
	Deviceid   int64 `json:"deviceid" binding:"required"`
}

type userActivityResponse struct {
	ID         int64     `json:"id"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Userid     int64     `json:"userid"`
	Deviceid   int64     `json:"deviceid"`
	Locationid int64     `json:"locationid"`
}

type DailyActivity struct {
	Day  int     `json:"day"`
	Time float64 `json:"time"`
}

type dailyUserActivityResponse struct {
	Month    int             `json:"month"`
	Activity []DailyActivity `json:"activity"`
}

func newUserActivityResponse(activity db.Useractivity) userActivityResponse {
	return userActivityResponse{
		ID:         activity.ID,
		Start:      activity.Start,
		End:        activity.End,
		Userid:     activity.Userid,
		Deviceid:   activity.Deviceid,
		Locationid: activity.Locationid,
	}
}

// CreateTags		godoc
// @Summary			Create CheckinActivity
// @Description 	Create CheckinActivity data in Db.
// @Param 			device body createUserActivityRecordRequest true "Create Checkin Activity Record"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} string
// @Router			/startActivity [post]
func (server *Server) createStartActicityRecord(ctx *gin.Context) {
	var req createUserActivityRecordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateActivityRecordsParams{
		Userid:     req.Userid,
		Type:       req.Type,
		Locationid: req.Locationid,
	}

	_, err := server.store.CreateActivityRecords(ctx, arg)
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

	ctx.JSON(http.StatusOK, "Start Activity Recorded")
}

// CreateTags		godoc
// @Summary			Create CheckinActivity
// @Description 	Create CheckinActivity data in Db.
// @Param 			device body createUserActivityRecordRequest true "Create Checkin Activity Record"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} userActivityResponse{}
// @Router			/endActivity [post]
func (server *Server) createEndActivityRecord(ctx *gin.Context) {
	var req createUserActivityRecordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateActivityRecordsParams{
		Userid:     req.Userid,
		Type:       req.Type,
		Locationid: req.Locationid,
		Deviceid:   req.Deviceid,
	}

	activity, err := service.CreateEndActivityRecord(ctx, server.store, arg)
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
	rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Create UserActivity
// @Description 	Create UserActivity data in Db.
// @Param 			device body createUserActivityRequest true "Create Device"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} userActivityResponse{}
// @Router			/userActivity [post]
func (server *Server) createUserActivity(ctx *gin.Context) {
	var req createUserActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateUserActivityParams{
		Start:    service.GetFormatedTime(req.Start),
		End:      service.GetFormatedTime(req.End),
		Userid:   req.Userid,
		Deviceid: req.Deviceid,
	}

	fmt.Println(arg)

	activity, err := server.store.CreateUserActivity(ctx, arg)
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

	rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get User Activity From ID
// @Description 	Get User Activity data from Db.
// @Param 			users query getUserActivityRequest true "Get user"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} []userActivityResponse{}
// @Router			/userActivity [get]
func (server *Server) getUserActivity(ctx *gin.Context) {
	var req getUserActivityRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	activity, err := server.store.GetUserActivity(ctx, req.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, activity)
}

// CreateTags		godoc
// @Summary			Get User Activity From ID
// @Description 	Get User Activity data from Db.
// @Param 			users query getPastWorkOutActivityRequest true "Get user"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} []db.GetPastWorkoutDataRow{}
// @Router			/getPastWorkoutData [get]
func (server *Server) getPastWorkoutData(ctx *gin.Context) {
	var req getPastWorkOutActivityRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	activity := []db.GetPastWorkoutDataRow{}
	err := errors.New("Switch")
	switch req.Interval {
	case "1":
		{
			activity, err = server.store.GetPastWorkoutData1(ctx, req.Userid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	case "7":
		{
			activity, err = server.store.GetPastWorkoutData7(ctx, req.Userid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	case "30":
		{
			activity, err = server.store.GetPastWorkoutData30(ctx, req.Userid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	case "60":
		{
			activity, err = server.store.GetPastWorkoutData60(ctx, req.Userid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	case "90":
		{
			activity, err = server.store.GetPastWorkoutData90(ctx, req.Userid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}

	}

	//rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, activity)
}

// CreateTags		godoc
// @Summary			Get Day wise User Activity From ID
// @Description 	Get Day Wise User Activity data from Db.
// @Param 			users query getUserActivityRequest true "Get user"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} []api.dailyUserActivityResponse{}
// @Router			/getDayWiseUserActivity [get]
func (server *Server) getDayWiseUserActivity(ctx *gin.Context) {
	var req getUserActivityRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	activity, err := server.store.GetDayWiseActivity(ctx, req.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rsp []dailyUserActivityResponse
	for i := 0; i < 12; i += 1 {
		monthly := dailyUserActivityResponse{}
		var daily []DailyActivity
		monthly.Month = i + 1
		days := daysInMonth(time.Now(), time.Month(i+1))
		for j := 0; j < days; j += 1 {
			curr := DailyActivity{
				Day:  j + 1,
				Time: 0,
			}
			daily = append(daily, curr)
		}
		monthly.Activity = daily
		rsp = append(rsp, monthly)
	}
	for _, a := range activity {
		month := a.Date.Month()
		day := a.Date.Day()
		fmt.Println("Date : ", a.Date)
		fmt.Println("month   : ", month, "  day   : ", day)
		rsp[month-1].Activity[day-1].Time = a.TotalTimeSeconds

	}

	//rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, rsp)
}

func daysInMonth(t time.Time, month time.Month) int {
	t = time.Date(t.Year(), month, 32, 0, 0, 0, 0, time.UTC)
	daysInMonth := 32 - t.Day()

	return daysInMonth
}
