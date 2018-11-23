package defs

type UserCredential struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	Sessioned string `json:"session_id"`
}

//时间都是按mysql得current_time设计
type VideoInfo struct {
	Id          int
	AuthorId    int
	Name        string
	Displaytime int
}

type Comments struct {
	Id       int
	AuthorId int
	VideoId  int
	Content  string
}
