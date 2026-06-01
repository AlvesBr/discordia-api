package model

import "time"

type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

type PaginatedPosts struct {
	Data []Post         `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

type Post struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Handle         string    `json:"handle"`
	Initials       string    `json:"initials"`
	AvatarGradient string    `json:"avatarGradient"`
	Body           string    `json:"body"`
	Likes          int       `json:"likes"`
	Reposts        int       `json:"reposts"`
	Liked          bool      `json:"liked"`
	Reposted       bool      `json:"reposted"`
	CommentCount   int       `json:"commentCount"`
	CreatedAt      time.Time `json:"createdAt"`
}
