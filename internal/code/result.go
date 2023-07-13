package code

type FinalResponse struct {
	Code    int         `json:"code"`    // 业务码
	Message string      `json:"message"` // 描述信息
	Body    interface{} `json:"body"`    // 数据
}
