package repository

import (
	"api-go/model"
	"database/sql"
	"fmt"
)

type CommentRepository struct {
	connection *sql.DB
}

func NewCommentRepository(conn *sql.DB) CommentRepository {
	return CommentRepository{connection: conn}
}

func (cr *CommentRepository) GetCommentsByPostId(postId int) ([]model.Comment, error) {
	query := "SELECT id, post_id, name, handle, initials, avatar_gradient, body, created_at FROM comments WHERE post_id = $1 ORDER BY created_at ASC"
	rows, err := cr.connection.Query(query, postId)
	if err != nil {
		fmt.Println(err)
		return []model.Comment{}, err
	}
	defer rows.Close()

	commentList := []model.Comment{}
	for rows.Next() {
		var c model.Comment
		err := rows.Scan(&c.ID, &c.PostID, &c.Name, &c.Handle, &c.Initials, &c.AvatarGradient, &c.Body, &c.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return []model.Comment{}, err
		}
		commentList = append(commentList, c)
	}

	return commentList, nil
}

func (cr *CommentRepository) CreateComment(comment model.Comment) (model.Comment, error) {
	query, err := cr.connection.Prepare(
		"INSERT INTO comments (post_id, name, handle, initials, avatar_gradient, body) " +
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at",
	)
	if err != nil {
		fmt.Println(err)
		return model.Comment{}, err
	}
	defer query.Close()

	err = query.QueryRow(
		comment.PostID,
		comment.Name,
		comment.Handle,
		comment.Initials,
		comment.AvatarGradient,
		comment.Body,
	).Scan(&comment.ID, &comment.CreatedAt)

	if err != nil {
		fmt.Println(err)
		return model.Comment{}, err
	}

	return comment, nil
}
