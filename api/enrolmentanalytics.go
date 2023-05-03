package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getOnLocationIdRequest struct {
	Locationid int64 `form:"locationid" binding:"required"`
}

// CreateTags		godoc
// @Summary			get ClassesOfferedAndAttendees
// @Description 	get ClassesOfferedAndAttendees data in Db.
// @Param 			ByLocation query getOnLocationIdRequest true "get ClassesOfferedAndAttendeesRequest"
// @Produce 		application/json
// @Tags 			enrolmentanalytics
// @Success 		200 {object} []db.GetClassesOfferedAndAttendeesRow{}
// @Router			/classesOfferedAndAttendes [get]
func (server *Server) getClassesOfferedAndAttendees(ctx *gin.Context) {
	var req getOnLocationIdRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	classes, err := server.store.GetClassesOfferedAndAttendees(ctx, req.Locationid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, classes)
}

// CreateTags		godoc
// @Summary			get AllClassesOfferedAndAttendees
// @Description 	get AllClassesOfferedAndAttendees data in Db.
// @Produce 		application/json
// @Tags 			enrolmentanalytics
// @Success 		200 {object} []db.GetAllClassesOfferedAndAttendeesRow{}
// @Router			/allClassesOfferedAndAttendes [get]
func (server *Server) getAllClassesOfferedAndAttendees(ctx *gin.Context) {

	data, err := server.store.GetAllClassesOfferedAndAttendees(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get ClassesOfferedAndAttendees
// @Description 	get ClassesOfferedAndAttendees data in Db.
// @Param 			ByLocation query getOnLocationIdRequest true "get ClassesOfferedAndAttendeesRequest"
// @Produce 		application/json
// @Tags 			enrolmentanalytics
// @Success 		200 {object} []db.GetClassesOfferedAndAttendeesPerWeekRow{}
// @Router			/classesOfferedAndAttendesPerWeek [get]
func (server *Server) getClassesOfferedAndAttendeesPerWeek(ctx *gin.Context) {
	var req getOnLocationIdRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//TODO :: Convert date_trunc in GetClassesOfferedAndAttendeesPerWeekRow to time.Time
	classes, err := server.store.GetClassesOfferedAndAttendeesPerWeek(ctx, req.Locationid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, classes)
}

// CreateTags		godoc
// @Summary			get AllClassesOfferedAndAttendeesPerWeek
// @Description 	get AllClassesOfferedAndAttendeesPerWeek data in Db.
// @Produce 		application/json
// @Tags 			enrolmentanalytics
// @Success 		200 {object} []db.GetAllClassesOfferedAndAttendeesPerWeekRow{}
// @Router			/allClassesOfferedAndAttendesPerWeek [get]
func (server *Server) getAllClassesOfferedAndAttendeesPerWeek(ctx *gin.Context) {

	//TODO :: Convert date_trunc in GetAllClassesOfferedAndAttendeesPerWeekRow to time.Time
	data, err := server.store.GetAllClassesOfferedAndAttendeesPerWeek(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get AllTopAttendedClass
// @Description 	get AllTopAttendedClass data in Db.
// @Produce 		application/json
// @Tags 			enrolmentanalytics
// @Success 		200 {object} []db.GetAllTopAttendedClassRow{}
// @Router			/allTopAttendedClass [get]
func (server *Server) getAllTopAttendedClass(ctx *gin.Context) {

	data, err := server.store.GetAllTopAttendedClass(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get MostPopularHourForClassesOnWeekdays
// @Description 	get MostPopularHourForClassesOnWeekdays data in Db.
// @Produce 		application/json
// @Tags 			enrolmentanalytics
// @Success 		200 {object} []db.GetMostPopularHourForClassesOnWeekdaysRow{}
// @Router			/mostPopularHourForClassesOnWeekdays [get]
func (server *Server) getMostPopularHourForClassesOnWeekdays(ctx *gin.Context) {

	data, err := server.store.GetMostPopularHourForClassesOnWeekdays(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get MostPopularHourForClassesOnWeekends
// @Description 	get MostPopularHourForClassesOnWeekends data in Db.
// @Produce 		application/json
// @Tags 			enrolmentAnalytics
// @Success 		200 {object} []db.GetMostPopularHourForClassesOnWeekendsRow{}
// @Router			/mostPopularHourForClassesOnWeekends [get]
func (server *Server) getMostPopularHourForClassesOnWeekends(ctx *gin.Context) {

	data, err := server.store.GetMostPopularHourForClassesOnWeekends(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}
