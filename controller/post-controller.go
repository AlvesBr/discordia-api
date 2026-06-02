package controller

import (
	"api-go/model"
	"api-go/repository"
	"api-go/usecase"
	"net/http"
	"strconv"
	"time"

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

	tag := ctx.DefaultQuery("tag", "")
	if !repository.IsSafeTag(tag) {
		tag = ""
	}

	result, err := p.postUseCase.GetPosts(page, limit, tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *postController) GetPostsSince(ctx *gin.Context) {
	tsStr := ctx.Query("ts")
	if tsStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "parâmetro ts é obrigatório"})
		return
	}

	since, err := time.Parse(time.RFC3339Nano, tsStr)
	if err != nil {
		since, err = time.Parse(time.RFC3339, tsStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ts inválido, use formato RFC3339"})
			return
		}
	}

	posts, err := p.postUseCase.GetPostsSince(since)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if posts == nil {
		posts = []model.Post{}
	}
	ctx.JSON(http.StatusOK, posts)
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
