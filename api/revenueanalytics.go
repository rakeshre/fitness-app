package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTags		godoc
// @Summary			get ClassRevenueGenerateByLocation
// @Description 	get ClassRevenueGenerateByLocation data in Db.
// @Produce 		application/json
// @Tags 			revenueanalytics
// @Success 		200 {object} []db.GetClassRevenueGenerateByLocationRow{}
// @Router			/classRevenueGeneratedByLocation [get]
func (server *Server) getClassRevenueGenerateByLocation(ctx *gin.Context) {
	//TODO:: GetHoursSpentInGymByDayRow  time.Time
	data, err := server.store.GetClassRevenueGenerateByLocation(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get RevenueGenerateByMemberships
// @Description 	get RevenueGenerateByMemberships data in Db.
// @Produce 		application/json
// @Tags 			revenueanalytics
// @Success 		200 {object} int64
// @Router			/revenueGenerateByMemberships [get]
func (server *Server) getRevenueGenerateByMemberships(ctx *gin.Context) {

	data, err := server.store.GetRevenueGenerateByMemberships(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}
