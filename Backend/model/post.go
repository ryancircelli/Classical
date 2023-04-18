package model

type Post struct {
	PostID        int64  `json:"postId"`
	PostClassName string `json:"postClassName"`
	PostName      string `json:"postName"`
	PostContent   string `json:"postContent"`
	PostVotes     int64  `json:"postVotes"`
	TimePosted    Time   `json:"timePosted"`
}
