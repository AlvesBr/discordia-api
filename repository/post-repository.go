package repository

import (
	"api-go/model"
	"database/sql"
	"fmt"
)

type PostRepository struct {
	connection *sql.DB
}

func NewPostRepository(conn *sql.DB) PostRepository {
	return PostRepository{
		connection: conn,
	}
}

func (pr *PostRepository) GetPostsCount() (int, error) {
	var total int
	err := pr.connection.QueryRow("SELECT COUNT(*) FROM posts").Scan(&total)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return total, nil
}

func (pr *PostRepository) PostExists(id int) (bool, error) {
	var count int
	err := pr.connection.QueryRow("SELECT COUNT(*) FROM posts WHERE id = $1", id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (pr *PostRepository) GetPosts(limit, offset int) ([]model.Post, error) {
	query := `SELECT id, name, handle, initials, avatar_gradient, body, likes, reposts, liked, reposted, created_at,
		(SELECT COUNT(*) FROM comments WHERE post_id = posts.id) AS comment_count
		FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := pr.connection.Query(query, limit, offset)
	if err != nil {
		fmt.Println(err)
		return []model.Post{}, err
	}
	defer rows.Close()

	var postList []model.Post
	for rows.Next() {
		var postObj model.Post
		err := rows.Scan(
			&postObj.ID,
			&postObj.Name,
			&postObj.Handle,
			&postObj.Initials,
			&postObj.AvatarGradient,
			&postObj.Body,
			&postObj.Likes,
			&postObj.Reposts,
			&postObj.Liked,
			&postObj.Reposted,
			&postObj.CreatedAt,
			&postObj.CommentCount,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Post{}, err
		}
		postList = append(postList, postObj)
	}

	return postList, nil
}

func (pr *PostRepository) CreatePost(post model.Post) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO posts " +
		"(name, handle, initials, avatar_gradient, body, likes, reposts, liked, reposted) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(
		post.Name,
		post.Handle,
		post.Initials,
		post.AvatarGradient,
		post.Body,
		post.Likes,
		post.Reposts,
		post.Liked,
		post.Reposted,
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (pr *PostRepository) ToggleLike(id int) error {
	query := "UPDATE posts SET liked = NOT liked, likes = CASE WHEN NOT liked THEN likes + 1 ELSE CASE WHEN likes > 0 THEN likes - 1 ELSE 0 END END WHERE id = $1"
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (pr *PostRepository) ToggleRepost(id int) error {
	query := "UPDATE posts SET reposted = NOT reposted, reposts = CASE WHEN NOT reposted THEN reposts + 1 ELSE CASE WHEN reposts > 0 THEN reposts - 1 ELSE 0 END END WHERE id = $1"
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
