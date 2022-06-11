package status

type Reply struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// OK 响应成功
var OK = Reply{Code: 0, Message: "SUCCESS"}
