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

type createClassRequest struct {
	Instructorname string `json:"instructorname" binding:"required"`
	Startdate      string `json:"startdate" binding:"required"`
	Enddate        string `json:"enddate" binding:"required"`
	Starttime      string `json:"starttime" binding:"required"`
	Endtime        string `json:"endtime" binding:"required"`
	Day            string `json:"day" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Locationid     int64  `json:"locationid" binding:"required"`
	Cost           int32  `json:"cost" binding:"required"`
}

type getClassFromIDRequest struct {
	Classid int64 `form:"classid" binding:"required"`
}

type getClassFromNameRequest struct {
	Classname string `form:"classname" binding:"required"`
}

type classResponse struct {
	ID             int64     `json:"id"`
	Instructorname string    `json:"instructorname"`
	Regstatus      string    `json:"regstatus"`
	Startdate      time.Time `json:"startdate"`
	Enddate        time.Time `json:"enddate"`
	Starttime      time.Time `json:"starttime"`
	Endtime        time.Time `json:"endtime"`
	Day            string    `json:"day"`
	Name           string    `json:"name"`
	// weekly daily or monthly
	Classtype  string `json:"classtype"`
	Locationid int64  `json:"locationid"`
	Cost       int32  `json:"cost"`
}

func newClassResponse(class db.Class, schedule db.Schedule) classResponse {
	return classResponse{
		ID:             class.ID,
		Instructorname: class.Instructorname,
		Regstatus:      class.Regstatus,
		Endtime:        schedule.Endtime,
		Starttime:      schedule.Starttime,
		Name:           class.Name,
		Classtype:      class.Classtype,
		Cost:           class.Cost,
		Day:            schedule.Day,
		Enddate:        schedule.Enddate,
		Startdate:      schedule.Startdate,
		Locationid:     schedule.Locationid,
	}
}

// CreateTags		godoc
// @Summary			Create Class
// @Description 	Create Class data in Db.
// @Param 			class body createClassRequest true "Create class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} classResponse{}
// @Router			/class [post]
func (server *Server) createClass(ctx *gin.Context) {
	var req createClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Invalidadata")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateScheduleParams{
		Endtime:    service.GetFormatedTime(req.Endtime),
		Starttime:  service.GetFormatedTime(req.Starttime),
		Day:        req.Day,
		Enddate:    service.GetFormatedDate(req.Enddate),
		Startdate:  service.GetFormatedDate(req.Startdate),
		Locationid: req.Locationid,
	}

	schedule, err := server.store.CreateSchedule(ctx, arg)
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

	arg1 := db.CreateClassParams{
		Instructorname: req.Instructorname,
		Name:           req.Name,
		Cost:           req.Cost,
		Scheduleid:     schedule.ID,
	}

	class, err := server.store.CreateClass(ctx, arg1)
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

	rsp := newClassResponse(class, schedule)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			get Class
// @Description 	get Class data in Db.
// @Param 			class query getClassFromIDRequest true "get class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} userResponse{}
// @Router			/class [get]
func (server *Server) getClass(ctx *gin.Context) {
	var req getClassFromIDRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	class, err := server.store.GetClass(ctx, req.Classid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	schedule, err := server.store.GetSchedule(ctx, class.Scheduleid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newClassResponse(class, schedule)
	ctx.JSON(http.StatusOK, rsp)
}

type getClassRequest struct {
	Locationid int64  `form:"locationid" binding:"required"`
	Day        string `form:"day" binding:"required"`
	Userid     int64  `form:"userid" binding:"required"`
}

// CreateTags		godoc
// @Summary			get Class
// @Description 	get Class data in Db.
// @Param 			class query getClassRequest true "get class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} []db.GetClassesRow{}
// @Router			/getClasses [get]
func (server *Server) getClasses(ctx *gin.Context) {
	var req getClassRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetClassesParams{
		Userid:     req.Userid,
		Locationid: req.Locationid,
		Day:        req.Day,
	}
	class, err := server.store.GetClasses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, class)
}

type getClassForEmployeeRequest struct {
	Day        string `form:"day" binding:"required"`
	Locationid int64  `form:"locationid" binding:"required"`
}

// CreateTags		godoc
// @Summary			get Class
// @Description 	get Class data in Db.
// @Param 			class query getClassForEmployeeRequest true "get class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} []db.GetClassesForEmployeeRow{}
// @Router			/getClassesForEmployee [get]
func (server *Server) getClassesForEmployee(ctx *gin.Context) {
	var req getClassForEmployeeRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetClassesForEmployeeParams{
		Locationid: req.Locationid,
		Day:        req.Day,
	}

	class, err := server.store.GetClassesForEmployee(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, class)
}

type getUpcomingClassesRequest struct {
	Userid int64 `form:"userid" binding:"required"`
}

// CreateTags		godoc
// @Summary			get Class
// @Description 	get Class data in Db.
// @Param 			class query getUpcomingClassesRequest true "get class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} []db.GetUpcomingClassesRow{}
// @Router			/getUpcomingClasses [get]
func (server *Server) getUpcomingClasses(ctx *gin.Context) {
	var req getUpcomingClassesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	class, err := server.store.GetUpcomingClasses(ctx, req.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, class)
}
