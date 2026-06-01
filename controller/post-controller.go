package controller

import (
	"api-go/model"
	"api-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postController struct {
	postUseCase usecase.PostUseCase
}

func NewPostController(usecase usecase.PostUseCase) postController {
	return postController{
		postUseCase: usecase,
	}
}

func (p *postController) GetPosts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	result, err := p.postUseCase.GetPosts(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *postController) CreatePost(ctx *gin.Context) {
	var post model.Post
	err := ctx.BindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedPost, err := p.postUseCase.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedPost)
}

func (p *postController) ToggleLike(ctx *gin.Context) {
	postId, errResp := parsePostID(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	err := p.postUseCase.ToggleLike(postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Like alternado com sucesso"})
}

func (p *postController) ToggleRepost(ctx *gin.Context) {
	postId, errResp := parsePostID(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	err := p.postUseCase.ToggleRepost(postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Repost alternado com sucesso"})
}

func parsePostID(ctx *gin.Context) (int, *model.Response) {
	id := ctx.Param("id")
	if id == "" {
		return 0, &model.Response{Message: "Id do post é obrigatório"}
	}

	postId, err := strconv.Atoi(id)
	if err != nil {
		return 0, &model.Response{Message: "Id do post deve ser um número"}
	}

	return postId, nil
}
