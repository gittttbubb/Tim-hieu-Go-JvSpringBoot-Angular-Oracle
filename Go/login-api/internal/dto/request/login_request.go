package request
// Struct hứng dữ liệu JSON gửi lên từ Client
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}