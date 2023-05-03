package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createLocationRequest struct {
	City    string `json:"city" binding:"required" `
	State   string `json:"state" binding:"required"`
	Zipcode string `json:"zipcode" binding:"required"`
}

type getLocationRequest struct {
	Locationid int64 `form:"locationid" binding:"required"`
}

type locationResponse struct {
	ID      int64  `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode" `
}

func newLocationResponse(location db.Location) locationResponse {
	return locationResponse{
		ID:      location.ID,
		City:    location.City,
		State:   location.State,
		Zipcode: location.Zipcode,
	}
}

// CreateTags		godoc
// @Summary			Create Location
// @Description 	Create Location data in Db.
// @Param 			users body createLocationRequest true "Create Location"
// @Produce 		application/json
// @Tags 			location
// @Success 		200 {object} locationResponse{}
// @Router			/location [post]
func (server *Server) createLocation(ctx *gin.Context) {
	var req createLocationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateLocationParams{
		City:    req.City,
		State:   req.State,
		Zipcode: req.Zipcode,
	}

	location, err := server.store.CreateLocation(ctx, arg)
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

	rsp := newLocationResponse(location)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get Location From ID
// @Description 	Get User data from Db.
// @Param 			users query getLocationRequest true "Get user"
// @Produce 		application/json
// @Tags 			location
// @Success 		200 {object} locationResponse{}
// @Router			/location [get]
func (server *Server) getLocation(ctx *gin.Context) {
	var req getLocationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	location, err := server.store.GetLocation(ctx, req.Locationid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newLocationResponse(location)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get ALL Location From
// @Description 	Get ALL locations data from Db.
// @Produce 		application/json
// @Tags 			location
// @Success 		200 {object} []db.Location{}
// @Router			/alllocations [get]
func (server *Server) getAllLocations(ctx *gin.Context) {

	locations, err := server.store.GetAllLocations(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := newLocationResponse(location)
	ctx.JSON(http.StatusOK, locations)
}
