package entity

type PostReaction struct {
	ID       uint64   `json:"id,omitempty"`
	PostID   uint64   `json:"post_id,omitempty"`
	UserID   uint64   `json:"user_id,omitempty"`
	Reaction Reaction `json:"reaction,omitempty"`
}

type CommentReaction struct {
	ID        uint64   `json:"id,omitempty"`
	CommentID uint64   `json:"comment_id,omitempty"`
	UserID    uint64   `json:"user_id,omitempty"`
	Reaction  Reaction `json:"reaction,omitempty"`
}

type Reaction struct {
	Type  string `json:"type,omitempty"`
	State bool   `json:"state,omitempty"`
}

// 0- none
// 1- post
// 2- comment
