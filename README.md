Họ và tên: Ngô Phi Nhật
Mã sinh viên:
Lớp: SE-K23
Bài thi: Adapter Pattern

## Việc cần làm

Đọc `assignment.md` trong bộ đề. Hoàn thành các TODO trong `main.go`: tạo `SMSAdapter` chứa `*LegacySMS`, triển khai `Notifier.Send()` và dùng adapter qua biến kiểu `Notifier` trong `main()`. Không sửa thư viện `LegacySMS`.

Starter hiện biên dịch được nhưng chưa gửi thông báo. Không yêu cầu unit test, database hay HTTP API.

## Kiểm tra và chạy

Từ thư mục chứa `go.mod` (Go 1.22 trở lên):

```bash
go build ./...
go run .
```

Sau khi hoàn thành, kết quả cần là:

```text
SMS: Xin chao, don hang cua ban da duoc xac nhan!
```

## trả lời câu hỏi

Mỗi câu trả lời khoảng 2–3 dòng.

1. Adapter Pattern giải quyết vấn đề gì?

   Adapter Pattern giải quyết việc hai thành phần có giao diện không khớp nhau nên không dùng chung được. Ở đây ứng dụng cần `Notifier.Send()`, còn thư viện cũ chỉ có `SendSMS()`. Adapter đứng giữa, nhận lời gọi `Send()` rồi chuyển thành `SendSMS()`.

2. Tại sao không cần chỉnh sửa `LegacySMS`?

   Vì `SMSAdapter` bọc `*LegacySMS` bên trong và tự gọi `SendSMS()` khi cần, nên mọi việc chuyển đổi diễn ra trong adapter. `LegacySMS` giữ nguyên, không bị ảnh hưởng. Nhờ vậy code cũ không có nguy cơ lỗi, và ta vẫn dùng được thư viện nếu không có quyền sửa nó.
