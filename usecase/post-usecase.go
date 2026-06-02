package usecase

import (
	"api-go/model"
	"api-go/repository"
	"strings"
	"time"
)

type PostUseCase struct {
	repository repository.PostRepository
}

func NewPostUseCase(repository repository.PostRepository) PostUseCase {
	return PostUseCase{
		repository: repository,
	}
}

func (pu *PostUseCase) GetPosts(page, limit int, tag string) (model.PaginatedPosts, error) {
	offset := (page - 1) * limit

	posts, err := pu.repository.GetPosts(limit, offset, tag)
	if err != nil {
		return model.PaginatedPosts{}, err
	}

	total, err := pu.repository.GetPostsCount(tag)
	if err != nil {
		return model.PaginatedPosts{}, err
	}

	totalPages := (total + limit - 1) / limit
	if totalPages < 1 {
		totalPages = 1
	}

	return model.PaginatedPosts{
		Data: posts,
		Meta: model.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (pu *PostUseCase) GetPostsSince(since time.Time) ([]model.Post, error) {
	return pu.repository.GetPostsSince(since)
}

func (pu *PostUseCase) CreatePost(post model.Post) (model.Post, error) {
	if strings.TrimSpace(post.Name) == "" {
		post.Name = "Anônimo"
	}

	if strings.TrimSpace(post.Initials) == "" {
		post.Initials = generateInitials(post.Name)
	}

	if strings.TrimSpace(post.Handle) == "" {
		post.Handle = generateHandle(post.Name)
	}

	if strings.TrimSpace(post.AvatarGradient) == "" {
		post.AvatarGradient = generateAvatarGradient()
	}

	post.Likes = 0
	post.Reposts = 0
	post.Liked = false
	post.Reposted = false
	post.CreatedAt = time.Now()

	postId, err := pu.repository.CreatePost(post)
	if err != nil {
		return model.Post{}, err
	}

	post.ID = postId

	return post, nil
}

func (pu *PostUseCase) ToggleLike(id int) error {
	return pu.repository.ToggleLike(id)
}

func (pu *PostUseCase) ToggleRepost(id int) error {
	return pu.repository.ToggleRepost(id)
}
