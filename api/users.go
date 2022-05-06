package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/gitnyasha/go-hekani-backend/db/sqlc"
	"github.com/gitnyasha/go-hekani-backend/util"
)

type createUserRequest struct {
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	HashedPassword string    `json:"hashed_password" binding:"required,min=8"`
	Bio            string    `json:"bio" binding:"required"`
	Birth          time.Time `json:"birth" binding:"required"`
	Image          string    `json:"image" binding:"required"`
}

type userResponse struct {
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" binding:"required,email"`
	Bio   string    `json:"bio" binding:"required"`
	Birth time.Time `json:"birth" binding:"required"`
	Image string    `json:"image" binding:"required"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Name:  user.Name,
		Email: user.Email,
		Bio:   user.Bio,
		Birth: user.Birth,
		Image: user.Image,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Bio:            req.Bio,
		Birth:          req.Birth,
		Image:          req.Image,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (int32(req.PageID) - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// create empty array
	resp := make([]userResponse, 0)

	for _, user := range users {
		// add items into arry
		resp = append(resp, newUserResponse(user))
	}

	ctx.JSON(http.StatusOK, resp)

}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"hashed_password" binding:"required,min=8"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_key"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.TokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, resp)
}
