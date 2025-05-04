## go mod init là gì ?
- Là lệnh để tạo file go.mod, dùng để quản lý module, dependencies (thư viện bên ngoài) và phiên bản Go mà project sử dụng.

 - Giống như package.json trong Node.js hoặc requirements.txt trong Python.

##  go mod
- go mod tidy	Xoá dependencies không dùng và thêm thiếu
- go mod why	Giải thích tại sao một package được import
- go mod download	Tải toàn bộ dependencies
- go list -m all	Liệt kê toàn bộ module phụ thuộc

## go get
- Dùng go get để thêm thư viện bên ngoài