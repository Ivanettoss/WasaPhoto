package components

type Username struct {
	Username string `json:"username"`
}

type User struct {
	Id       int    `json:"Id"`
	Username string `json:"username"`
}

type Profile struct {
	Username    string `json:"username"`
	NPhoto      int    `json:"nphoto"`
	NFollowers  int    `json:"nfollower"`
	NFollowed   int    `json:"nfollowed"`
	NPost       int    `json:"npost"`
	FollowState bool   `json:"followstate"`
	BanState    bool   `json:"banstate"`
}

type Photo struct {
	Username       string `json:"username"`
	PhotoBytes     string `json:"photobytes"`
	IdPhoto        int    `json:"idphoto"`
	NLikes         int    `json:"nlikes"`
	NComments      int    `json:"ncomments"`
	UploadDataTime string `json:"uploaddatatime"`
}

type Comment struct {
	Username       string `json:"username"`
	IdPhoto        int    `json:"idphoto"`
	IdComment      int    `json:"idcomment"`
	UploadDataTime string `json:"uploaddatatime"`
	Text           string `json:"text"`
}

type Stream struct {
	Username  string  `json:"username"`
	PhotoList []Photo `json:"photolist"`
}

type Ulist struct {
	UsersList []Username `json:"userlist"`
}
