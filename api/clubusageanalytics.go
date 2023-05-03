package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTags		godoc
// @Summary			get ClassesOfferedAndAttendees
// @Description 	get ClassesOfferedAndAttendees data in Db.
// @Produce 		application/json
// @Tags 			clubusageanalytics
// @Success 		200 {object} []db.GetHoursSpentInGymByDayRow{}
// @Router			/hoursSpentInGymByDay [get]
func (server *Server) getHoursSpentInGymByDay(ctx *gin.Context) {
	//TODO:: GetHoursSpentInGymByDayRow  time.Time
	classes, err := server.store.GetHoursSpentInGymByDay(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// CreateTags		godoc
// @Summary			get BusiestTimeByHourAndDayOfWeek
// @Description 	get BusiestTimeByHourAndDayOfWeek data in Db.
// @Produce 		application/json
// @Tags 			clubusageanalytics
// @Success 		200 {object} []db.GetBusiestTimeByHourAndDayOfWeekRow{}
// @Router			/busiestTimeByHourAndDayOfWeek [get]
func (server *Server) getBusiestTimeByHourAndDayOfWeek(ctx *gin.Context) {

	data, err := server.store.GetBusiestTimeByHourAndDayOfWeek(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get AverageVisitorsPerHourWeekdays
// @Description 	get AverageVisitorsPerHourWeekdays data in Db.
// @Produce 		application/json
// @Tags 			clubusageanalytics
// @Success 		200 {object} []db.GetAverageVisitorsPerHourWeekdaysRow{}
// @Router			/averageVisitorsPerHourWeekdays [get]
func (server *Server) getAverageVisitorsPerHourWeekdays(ctx *gin.Context) {
	//TODO :: Convert date_trunc in GetAverageVisitorsPerHourWeekdaysRow to time.Time
	classes, err := server.store.GetAverageVisitorsPerHourWeekdays(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, classes)
}

// CreateTags		godoc
// @Summary			get AverageVisitorsPerHourWeekends
// @Description 	get AverageVisitorsPerHourWeekends data in Db.
// @Produce 		application/json
// @Tags 			clubusageanalytics
// @Success 		200 {object} []db.GetAverageVisitorsPerHourWeekendsRow{}
// @Router			/averageVisitorsPerHourWeekends [get]
func (server *Server) getAverageVisitorsPerHourWeekends(ctx *gin.Context) {
	//TODO :: Convert date_trunc in GetAverageVisitorsPerHourWeekdaysRow to time.Time
	classes, err := server.store.GetAverageVisitorsPerHourWeekends(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, classes)
}
