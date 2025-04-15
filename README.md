# 🍔 ServerFastFood – Backend cho hệ thống quản lý cửa hàng thức ăn nhanh

Đây là phần backend của dự án **FastFood**, được xây dựng bằng ngôn ngữ **Go**. Dự án cung cấp API RESTful phục vụ cho hệ thống bán hàng thức ăn nhanh, bao gồm các chức năng như xác thực người dùng, quản lý sản phẩm, đơn hàng và vai trò người dùng. Phần frontend tương ứng có thể được tìm thấy tại [FastFoodTs](https://github.com/votadung2000/FastFoodTs).

## 🌐 Liên kết hệ thống

- Frontend: [FastFoodTs (React + TS)](https://github.com/votadung2000/FastFoodTs)
- Backend: [ServerFastFood (Go)](https://github.com/votadung2000/ServerFastFood)

## 🚀 Tính năng chính

- Đăng nhập, đăng ký và xác thực người dùng bằng JWT.
- Quản lý sản phẩm, danh mục, đơn hàng và trạng thái đơn hàng.
- Phân quyền người dùng theo vai trò (admin, khách hàng, v.v.).
- Hỗ trợ upload ảnh sản phẩm.
- Kết nối cơ sở dữ liệu bằng MySQL.
- Tích hợp Docker & Docker Compose để triển khai nhanh.

## 🛠️ Cài đặt và chạy dự án

### Yêu cầu

- Go >= 1.20
- Docker và Docker Compose
- MySQL

## 📁 Cấu trúc thư mục

```bash
ServerFastFood/
├── components/         # Các thành phần hỗ trợ như token, email, mã hóa,...
│   └── tokenProvider/  # Xử lý JWT và xác thực
├── database/           # Kết nối và thao tác với cơ sở dữ liệu (MySQL)
├── middleware/         # Các middleware như xác thực, logging, xử lý lỗi,...
├── modules/            # Business logic cho từng đối tượng: user, product, order,...
│   └── user/           # Ví dụ: xử lý logic liên quan đến người dùng
├── router/             # Định nghĩa các route và ánh xạ đến handler
├── static/             # Chứa tài nguyên tĩnh như ảnh sản phẩm,...
├── utils/              # Các hàm tiện ích dùng chung
├── main.go             # Điểm khởi đầu của ứng dụng
├── go.mod              # Thông tin module Go
├── .env                # Biến môi trường (dùng khi không chạy Docker)
└── docker-compose.yml  # Cấu hình Docker Compose
```


## 📡 API Endpoint mẫu

- POST /register – Đăng ký người dùng
- POST /login – Đăng nhập và nhận token
- GET /products – Lấy danh sách sản phẩm
- POST /orders – Tạo đơn hàng
  
👉 Tham khảo frontend để biết cách sử dụng API cụ thể: FastFoodTs

## 🤝 Đóng góp
Mọi đóng góp đều được hoan nghênh!
Bạn có thể:

1. Fork repository.
2. Tạo nhánh mới từ main.
3. Commit và gửi pull request.
