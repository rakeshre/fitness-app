package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTags		godoc
// @Summary			get GetDailyNewMemberEnrolments
// @Description 	get GetDailyNewMemberEnrolments data in Db.
// @Produce 		application/json
// @Tags 			memberretentionanalytics
// @Success 		200 {object} []db.GetDailyNewMemberEnrolmentsRow{}
// @Router			/dailyNewMemberEnrolments [get]
func (server *Server) getDailyNewMemberEnrolments(ctx *gin.Context) {
	//TODO:: time.Time
	classes, err := server.store.GetDailyNewMemberEnrolments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// CreateTags		godoc
// @Summary			get GetMembershipCountsByType
// @Description 	get GetMembershipCountsByType data in Db.
// @Produce 		application/json
// @Tags 			memberretentionanalytics
// @Success 		200 {object} []db.GetMembershipCountsByTypeRow{}
// @Router			/membershipCountsByType [get]
func (server *Server) getMembershipCountsByType(ctx *gin.Context) {

	data, err := server.store.GetMembershipCountsByType(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// CreateTags		godoc
// @Summary			get GetKMostFrequentMembers
// @Description 	get GetKMostFrequentMembers data in Db.
// @Produce 		application/json
// @Tags 			memberretentionanalytics
// @Success 		200 {object} []db.GetKMostFrequentMembersRow{}
// @Router			/kMostFrequentMembers [get]
func (server *Server) getKMostFrequentMembers(ctx *gin.Context) {
	classes, err := server.store.GetKMostFrequentMembers(ctx, 5)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, classes)
}
