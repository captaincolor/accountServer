package user

// 消息体有JSON参数需要传递，那么就针对每一个API接口定义独立的struct来接收
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}
