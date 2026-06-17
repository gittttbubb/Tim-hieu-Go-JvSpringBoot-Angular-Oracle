package utils

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.NewString()
}
// biến một giá trị string thành con trỏ tới string để có thể nil hoặc là 1 string
func StringPtr(s string) *string {
	return &s
}