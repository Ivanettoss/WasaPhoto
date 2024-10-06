package components

type Username struct {
	Username string `json:"username"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Profile struct {
	Username    string  `json:"username"`
	Photos      []Photo `json:"photos"`
	NFollowers  int     `json:"nfollower"`
	NFollowed   int     `json:"nfollowed"`
	NPost       int     `json:"npost"`
	FollowState bool    `json:"followstate"`
	BanState    bool    `json:"banstate"`
}

type Photo struct {
	Username       string `json:"username"`
	PhotoBytes     string `json:"photobytes"`
	IdPhoto        int    `json:"idphoto"`
	NLikes         int    `json:"nlikes"`
	NComments      int    `json:"ncomments"`
	UploadDataTime string `json:"uploaddatatime"`
	IsLiked        bool   `json:"isliked"`
}

type Comment struct {
	User           User   `json:"user"`
	IdPhoto        int    `json:"idphoto"`
	IdComment      int    `json:"idcomment"`
	UploadDataTime string `json:"uploaddatatime"`
	Text           string `json:"text"`
}

type Stream struct {
	Username string  `json:"username"`
	Photos   []Photo `json:"photos"`
}

type Ulist struct {
	UsersList []Username `json:"userlist"`
}
