package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createDeviceRequest struct {
	Description string `json:"description" binding:"required"`
	// Free,busy,not working
	Status string `json:"status"`
}

type getDeviceRequest struct {
	Deviceid int64 `form:"deviceid" binding:"required"`
}

type deviceResponse struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	// Free,busy,not working
	Status string `json:"status"`
}

func newDeviceResponse(device db.Device) deviceResponse {
	return deviceResponse{
		ID:          device.ID,
		Description: device.Description,
		Status:      device.Status,
	}
}

// CreateTags		godoc
// @Summary			Create Device
// @Description 	Create Device data in Db.
// @Param 			device body createDeviceRequest true "Create Device"
// @Produce 		application/json
// @Tags 			device
// @Success 		200 {object} deviceResponse{}
// @Router			/device [post]
func (server *Server) createDevice(ctx *gin.Context) {
	var req createDeviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateDeviceParams{
		Description: req.Description,
		Status:      req.Status,
	}

	device, err := server.store.CreateDevice(ctx, arg)
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

	rsp := newDeviceResponse(device)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get Device From ID
// @Description 	Get Device data from Db.
// @Param 			users query getDeviceRequest true "Get user"
// @Produce 		application/json
// @Tags 			device
// @Success 		200 {object} deviceResponse{}
// @Router			/device [get]
func (server *Server) getDevice(ctx *gin.Context) {
	var req getDeviceRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	device, err := server.store.GetDevice(ctx, req.Deviceid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newDeviceResponse(device)
	ctx.JSON(http.StatusOK, rsp)
}
