package controller

import (
	"api-go/model"
	"api-go/usecase"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentUseCase usecase.CommentUseCase
}

func NewCommentController(usecase usecase.CommentUseCase) commentController {
	return commentController{commentUseCase: usecase}
}

func (c *commentController) GetComments(ctx *gin.Context) {
	postId, err := parseCommentPostID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}

	comments, err := c.commentUseCase.GetCommentsByPostId(postId)
	if err != nil {
		if err.Error() == "post not found" {
			ctx.JSON(http.StatusNotFound, model.Response{Message: "Post não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *commentController) CreateComment(ctx *gin.Context) {
	postId, err := parseCommentPostID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}

	var comment model.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido"})
		return
	}

	created, err := c.commentUseCase.CreateComment(postId, comment)
	if err != nil {
		if err.Error() == "post not found" {
			ctx.JSON(http.StatusNotFound, model.Response{Message: "Post não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

func parseCommentPostID(ctx *gin.Context) (int, error) {
	id := ctx.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil || postId < 1 {
		return 0, errors.New("id inválido")
	}
	return postId, nil
}
