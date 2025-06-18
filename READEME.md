## Các loại print trong golang

```go
fmt.Print("Hello") // => Không xuống dòng
fmt.Println("Hello", "world!")  // Output: Hello world!\n => Tự động xuống dòng
name := "Go"
age := 15
fmt.Printf("Language: %s, Age: %d\n", name, age) // => In với định dạng (format) tương tự như trong C, Java.
// %s	Chuỗi (string)	"Go" → Go
// %d	Số nguyên (int)	10 → 10
// %f	Số thực (float)	3.14 → 3.140000
// %v	In giá trị mặc định	true → true, map → {}
// %+v	In giá trị và tên trường (struct)	{Name: "A"} → Name:A
// %T	In kiểu dữ liệu	10 → int
```

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
- Không thể mở rộng

```go
arr1 := [3]int{1, 2, 3}
arr2 := arr1       // sao chép toàn bộ giá trị

arr2[0] = 100

fmt.Println("arr1:", arr1) // [1 2 3] — không bị ảnh hưởng
fmt.Println("arr2:", arr2) // [100 2 3]
```

## Slice

- Kích thước linh hoạt, có thể thay đổi bằng append()
- Là kiểu tham chiếu → khi gán cho biến khác sẽ cùng trỏ về một mảng bên dưới (backing array)
- Thường được dùng trong thực tế vì tiện lợi
- Vì slice chỉ là view lên mảng, nên thay đổi slice sẽ thay đổi array gốc

```go
s1 := []int{1, 2, 3}
s2 := s1        // chỉ copy con trỏ, không copy dữ liệu
s2[0] = 100

fmt.Println("s1:", s1) // [100 2 3]
fmt.Println("s2:", s2) // [100 2 3]
```

## range

- Trong Golang, range là một từ khóa dùng để duyệt qua (iterate) các phần tử trong các kiểu dữ liệu có thể lặp như: array (mảng) , slice , map , string , channel

```go
import (
	"fmt"
)

for i, itemHightScore := range highScores {
	fmt.Println("Index:", i, "itemHightScore:", itemHightScore)
}
```

## Copy

- Trong Go (Golang), hàm copy() được sử dụng để sao chép các phần tử từ một slice vào slice khác.
- Nguyên tắc hoạt động
  1.copy chỉ hoạt động với slice.
  2.Nó sẽ sao chép tối đa min(len(dst), len(src)) phần tử.

```go
// copy(dst, src)
// dst: slice đích.
// src: slice nguồn.
// Kết quả trả về: số phần tử được sao chép thành công (int).

package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, 2)

	n := copy(dst, src)

	fmt.Println("dst:", dst)    // [1 2]
	fmt.Println("số phần tử đã copy:", n) // 2
}
```

## Pointers

- ❗ Lưu ý quan trọng: Không được dereference con trỏ nil

```go

package main

import (
	"fmt"
)

func main() {
	// - Thì mặc định y có giá trị là nil (tức là chưa trỏ đến vùng nhớ nào), nếu bạn cố *y = 10 thì chương trình sẽ panic vì dereference nil pointer.
  	var y *int

	// CÁCH 1:
	// a := 5
	// var y *int = &a
	// fmt.Println(*y) // 👉 5


	// Cách 2: Dùng new(int) để cấp phát vùng nhớ
	// var y *int = new(int)
	// *y = 10
	// fmt.Println(*y) // 👉 10

	// CÁCH 3 :  Khai báo nhanh bằng := và new()
	y := new(int)
	*y = 20
	fmt.Println(*y) // 👉 20

	age := 23
	// fmt.Println(*age) // ❌ lỗi: cannot indirect age (type int) => Vì *age nghĩa là: lấy giá trị tại địa chỉ mà age trỏ tới. Nhưng age chỉ là một giá trị số, không phải địa chỉ (con trỏ).

}
```
