package model

type Post struct {
	PostID      int64  `json:"postId"`
	ClassID     int64  `json:"classId"`
	PostName    string `json:"postName"`
	PostContent string `json:"postContent"`
	PostVotes   int64  `json:"postVotes"`
}
