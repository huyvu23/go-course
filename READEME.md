## go mod init là gì ?

- Là lệnh để tạo file go.mod, dùng để quản lý module, dependencies (thư viện bên ngoài) và phiên bản Go mà project sử dụng.

- Giống như package.json trong Node.js hoặc requirements.txt trong Python.

## go mod

- go mod tidy Xoá dependencies không dùng và thêm thiếu
- go mod why Giải thích tại sao một package được import
- go mod download Tải toàn bộ dependencies
- go list -m all Liệt kê toàn bộ module phụ thuộc

## go get

- Dùng go get để thêm thư viện bên ngoài

## Array

- Kích thước cố định, xác định tại thời điểm khai báo
- Là kiểu giá trị → khi gán cho biến khác sẽ copy toàn bộ
<!-- arr1 := [3]int{1, 2, 3}
arr2 := arr1       // sao chép toàn bộ giá trị

arr2[0] = 100

fmt.Println("arr1:", arr1) // [1 2 3] — không bị ảnh hưởng
fmt.Println("arr2:", arr2) // [100 2 3] -->

- Không thể mở rộng

## Slice

- Kích thước linh hoạt, có thể thay đổi bằng append()
- Là kiểu tham chiếu → khi gán cho biến khác sẽ cùng trỏ về một mảng bên dưới (backing array)
- Thường được dùng trong thực tế vì tiện lợi
<!-- s1 := []int{1, 2, 3}
s2 := s1        // chỉ copy con trỏ, không copy dữ liệu
s2[0] = 100

fmt.Println("s1:", s1) // [100 2 3]
fmt.Println("s2:", s2) // [100 2 3] -->
