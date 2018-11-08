package def

//request 直接提取url中json:<变量名>的value
type UserCredential struct{
	Username string 'json:"user_name"'
	Pwd string 'json:"pwd"'
}

//video video信息
type VideoInfo struct{
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

//comment 评论信息
type Comment struct{
	Id string
	VideoId string
	Author string
	Content string
}

//response 
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}