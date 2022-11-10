package responses

type UserResponse struct {
	Success bool `json:"success"`
	Message string `json:"msg"`
	Data map[string]interface{} `json:"data"`
}