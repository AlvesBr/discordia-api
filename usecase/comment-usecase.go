package usecase

import (
	"api-go/model"
	"api-go/repository"
	"errors"
	"strings"
	"time"
)

type CommentUseCase struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository, postRepo repository.PostRepository) CommentUseCase {
	return CommentUseCase{commentRepo: commentRepo, postRepo: postRepo}
}

func (cu *CommentUseCase) GetCommentsByPostId(postId int) ([]model.Comment, error) {
	exists, err := cu.postRepo.PostExists(postId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("post not found")
	}

	return cu.commentRepo.GetCommentsByPostId(postId)
}

func (cu *CommentUseCase) CreateComment(postId int, comment model.Comment) (model.Comment, error) {
	exists, err := cu.postRepo.PostExists(postId)
	if err != nil {
		return model.Comment{}, err
	}
	if !exists {
		return model.Comment{}, errors.New("post not found")
	}

	if strings.TrimSpace(comment.Name) == "" {
		comment.Name = "Anônimo"
	}

	comment.PostID = postId
	comment.Initials = generateInitials(comment.Name)
	comment.Handle = generateHandle(comment.Name)
	comment.AvatarGradient = generateAvatarGradient()
	comment.CreatedAt = time.Now()

	return cu.commentRepo.CreateComment(comment)
}
