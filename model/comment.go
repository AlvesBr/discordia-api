package model

import "time"

type Comment struct {
	ID             int       `json:"id"`
	PostID         int       `json:"postId"`
	Name           string    `json:"name"`
	Handle         string    `json:"handle"`
	Initials       string    `json:"initials"`
	AvatarGradient string    `json:"avatarGradient"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"createdAt"`
}
