package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createClassCatalogueRequest struct {
	Userid   int64 `json:"userid" binding:"required"`
	Courseid int64 `json:"courseid" binding:"required"`
}

type getClassCatalogueRequest struct {
	Userid int64 `form:"userid" binding:"required"`
}

type classCatalogueResponse struct {
	ID       int64 `json:"id"`
	Userid   int64 `json:"userid"`
	Courseid int64 `json:"courseid"`
}

func newClassCatalogueResponse(catalogue db.Classcatalogue) classCatalogueResponse {
	return classCatalogueResponse{
		ID:       catalogue.ID,
		Userid:   catalogue.Userid,
		Courseid: catalogue.Courseid,
	}
}

// CreateTags		godoc
// @Summary			Create ClassCatalogue
// @Description 	Create ClassCatalogue data in Db.
// @Param 			ClassCatalogue body createClassCatalogueRequest true "Create ClassCatalogue Activity"
// @Produce 		application/json
// @Tags 			classCatalogue
// @Success 		200 {object} classCatalogueResponse{}
// @Router			/classCatalogue [post]
func (server *Server) createClassCatalogue(ctx *gin.Context) {
	var req createClassCatalogueRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req)

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateClassCatalogueParams{
		Userid:   req.Userid,
		Courseid: req.Courseid,
	}

	catalogue, err := server.store.CreateClassCatalogue(ctx, arg)
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

	rsp := newClassCatalogueResponse(catalogue)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			get ClassCatalogue
// @Description 	get ClassCatalogue data in Db.
// @Param 			ClassCatalogue query getClassCatalogueRequest true "get ClassCatalogue Activity"
// @Produce 		application/json
// @Tags 			classCatalogue
// @Success 		200 {object} []classCatalogueResponse{}
// @Router			/classCatalogue [get]
func (server *Server) getClassCatalogue(ctx *gin.Context) {
	var req getClassCatalogueRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	classes, err := server.store.GetUserClass(ctx, req.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, classes)
}
