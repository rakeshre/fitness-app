package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createEmployeeRequest struct {
	Employeename string `json:"username" binding:"required,alphanum"`
	Password     string `json:"password" binding:"required,min=6"`
	Email        string `json:"email" binding:"required,email"`
	Locationid   int64  `json:"locationid" binding:"required"`
}

type getEmployeeFromIDRequest struct {
	UserId int64 `form:"userid" binding:"required"`
}

type getEmployeeFromNameRequest struct {
	Employeename string `form:"employeename" binding:"required"`
}

type employeeResponse struct {
	Employeename string    `json:"employeename"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	Locationid   int64     `json:"locationid"`
}

func newEmployeeResponse(user db.Employee) employeeResponse {
	return employeeResponse{
		Employeename: user.Name,
		Email:        user.Email,
		Locationid:   user.Locationid,
		CreatedAt:    user.CreatedAt,
	}
}

// CreateTags		godoc
// @Summary			Create Employee
// @Description 	Create Employee data in Db.
// @Param 			employee body createEmployeeRequest true "Create employee"
// @Produce 		application/json
// @Tags 			employee
// @Success 		200 {object} employeeResponse{}
// @Router			/employee [post]
func (server *Server) createEmployee(ctx *gin.Context) {
	var req createEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	//arg := db.CreateUserParams{
	//	Name:           req.Username,
	//	Hashedpassword: req.Password,
	//	Email:          req.Email,
	//}

	arg := db.CreateEmployeeParams{
		Name:           req.Employeename,
		Email:          req.Email,
		Hashedpassword: req.Password,
		Locationid:     req.Locationid,
	}

	user, err := server.store.CreateEmployee(ctx, arg)
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

	rsp := newEmployeeResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get Employee From EmployeeName
// @Description 	Get Employee data from Db.
// @Param 			users query getEmployeeFromNameRequest true "Get Employee"
// @Produce 		application/json
// @Tags 			employee
// @Success 		200 {object} userResponse{}
// @Router			/employee [get]
func (server *Server) getEmployee(ctx *gin.Context) {
	var req getEmployeeFromNameRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//username := ctx.Query("username")
	employee, err := server.store.GetEmployee(ctx, req.Employeename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newEmployeeResponse(employee)
	ctx.JSON(http.StatusOK, rsp)
}

//type loginUserRequest struct {
//	Username string `json:"username" binding:"required,alphanum"`
//	Password string `json:"password" binding:"required,min=6"`
//}
//
//type loginUserResponse struct {
//	AccessToken string       `json:"access_token"`
//	User        userResponse `json:"user"`
//}

//
//func (server *Server) loginUser(ctx *gin.Context) {
//	var req loginUserRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	user, err := server.store.GetUser(ctx, req.Username)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	err = util.CheckPassword(req.Password, user.HashedPassword)
//	if err != nil {
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	accessToken, err := server.tokenMaker.CreateToken(
//		user.Username,
//		server.config.AccessTokenDuration,
//	)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	rsp := loginUserResponse{
//		AccessToken: accessToken,
//		User:        newUserResponse(user),
//	}
//	ctx.JSON(http.StatusOK, rsp)
//}
