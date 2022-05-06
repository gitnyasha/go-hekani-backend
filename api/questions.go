package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gitnyasha/go-hekani-backend/db/sqlc"
)

type createQuestionRequest struct {
	Title              string `json:"title" binding:"required"`
	UserID             int32  `json:"user_id" binding:"required"`
	QuestionCategoryID int32  `json:"question_category_id" binding:"required"`
}

func (server *Server) createQuestion(ctx *gin.Context) {
	var req createQuestionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateQuestionParams{
		Title:              req.Title,
		UserID:             req.UserID,
		QuestionCategoryID: req.QuestionCategoryID,
	}

	question, err := server.store.CreateQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}

type getQuestionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getQuestion(ctx *gin.Context) {
	var req getQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	question, err := server.store.GetQuestion(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}

type listQuestionRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_id" binding:"required,min=5,max=10`
}

func (server *Server) listQuestion(ctx *gin.Context) {
	var req listQuestionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListQuestionsParams{
		Limit:  req.PageSize,
		Offset: (int32(req.PageID) - 1) * req.PageSize,
	}

	questions, err := server.store.ListQuestions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, questions)
}
