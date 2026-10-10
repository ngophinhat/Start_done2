package main

import "fmt"

// Notifier là interface mà ứng dụng sử dụng để gửi thông báo.
type Notifier interface {
	Send(message string) error
}

// LegacySMS là thư viện cũ. Không chỉnh sửa struct và phương thức này.
type LegacySMS struct{}

func (s *LegacySMS) SendSMS(content string) {
	fmt.Println("SMS:", content)
}

// SMSAdapter bọc LegacySMS để nó dùng được như một Notifier.
type SMSAdapter struct {
	legacy *LegacySMS
}

// Send chuyển lời gọi Send(message) thành SendSMS(message) của thư viện cũ.
func (a *SMSAdapter) Send(message string) error {
	a.legacy.SendSMS(message)
	return nil // thư viện cũ không trả lỗi
}

func main() {
	legacy := &LegacySMS{}
	adapter := &SMSAdapter{legacy: legacy}

	var notifier Notifier = adapter
	notifier.Send("Xin chao, don hang cua ban da duoc xac nhan!")
}
