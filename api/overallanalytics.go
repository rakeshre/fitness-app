package api

import (
	"Gym-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getOverAllAnalyticsRequest struct {
	Type int `form:"type" binding:"required"`
}

// CreateTags		godoc
// @Summary			get ClassRevenueGenerateByLocation
// @Description 	get ClassRevenueGenerateByLocation data in Db.
// @Param 			class query getOverAllAnalyticsRequest true "get overAllAnalyticsRequest"
// @Produce 		application/json
// @Tags 			overallanalytics
// @Success 		200 {object} string{}
// @Router			/getOverallAnalytics [get]
func (server *Server) getOverallAnalytics(ctx *gin.Context) {
	var req getOverAllAnalyticsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data, err := service.GetOverallAnalytics(ctx, server.store, req.Type)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, data)
}
