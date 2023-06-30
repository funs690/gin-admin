package auth

// session 数据
type Session struct {
	Id        string                 `json:"id"`
	UserId    string                 `json:"userId"`
	UserName  string                 `json:"username"`
	Token     string                 `json:"token"`
	ExpiresAt int64                  `json:"exp"`
	CreateAt  int64                  `json:"create"`
	Client    string                 `json:"client"`
	Data      map[string]interface{} `json:"data"`
}
