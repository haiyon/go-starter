package dto

// SampleBody 示例报文
type SampleBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}
