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

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type createNewUserRequest struct {
	Username   string `json:"username" binding:"required,alphanum"`
	Password   string `json:"password" binding:"required,min=6"`
	Email      string `json:"email" binding:"required,email"`
	Membership int64  `json:"membership" binding:"required"`
	Istrail    *bool  `json:"istrail" binding:"required"`
}

type getUserFromIDRequest struct {
	UserId int64 `form:"userid" binding:"required"`
}

type getUserFromEmailRequest struct {
	UserId int64 `form:"userid" binding:"required"`
}

type getUserFromNameRequest struct {
	Username string `form:"username" binding:"required"`
}

type userResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// CreateTags		godoc
// @Summary			Create User
// @Description 	Create User data in Db.
// @Param 			users body createNewUserRequest true "Create user"
// @Produce 		application/json
// @Tags 			user
// @Success 		200 {object} userResponse{}
// @Router			/usersV2 [post]
func (server *Server) createUserV2(ctx *gin.Context) {
	var req createNewUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Invalid Data")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateUserParams{
		Name:           req.Username,
		Hashedpassword: req.Password,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
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

	membershipargs := db.CreateMembershipParams{
		Membershipid: req.Membership,
		Userid:       user.ID,
		Expirydate:   time.Now().AddDate(0, int(req.Membership), 0),
	}

	_, err = server.store.CreateMembership(ctx, membershipargs)

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

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Create User
// @Description 	Create User data in Db.
// @Param 			users body createUserRequest true "Create user"
// @Produce 		application/json
// @Tags 			user
// @Success 		200 {object} userResponse{}
// @Router			/users [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateUserParams{
		Name:           req.Username,
		Hashedpassword: req.Password,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
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

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get User From UserName
// @Description 	Get User data from Db.
// @Param 			users query getUserFromNameRequest true "Get user"
// @Produce 		application/json
// @Tags 			user
// @Success 		200 {object} userResponse{}
// @Router			/users [get]
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserFromNameRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//username := ctx.Query("username")
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginResponse struct {
	Type   string `json:"type"`
	Error  string `json:"error"`
	Object string `json:"object"`
}

// CreateTags		godoc
// @Summary			Validate login Request
// @Description 	Validate login Request and return corresponding object.
// @Param 			login body loginRequest true "Valid User/Employee"
// @Produce 		application/json
// @Tags 			auth
// @Success 		200 {object} loginResponse{}
// @Router			/login [post]
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	validate, objecttype, obj := service.ValidateUser(ctx, req.Email, req.Password, server.store)

	rsp := loginResponse{
		Type:   objecttype,
		Error:  validate,
		Object: obj,
	}
	fmt.Println(rsp)

	ctx.JSON(http.StatusOK, rsp)
}
