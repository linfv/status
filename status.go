package status

type Reply struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// OK εεΊζε
var OK = Reply{Code: 0, Message: "SUCCESS"}
