# 🛍️ MyFiber API - Final Project Rakamin Evermos Virtual Internship

A RESTful API Toko (Store) built with Go (Golang) and Fiber framework.

## 📌 Table of Contents

- [Features](https://github.com/nandarahmat/my-fiber-api?tab=readme-ov-file#-features)
- [Tech Stack](https://github.com/nandarahmat/my-fiber-api?tab=readme-ov-file#-tech-stack)
- [API Endpoints](https://github.com/nandarahmat/my-fiber-api?tab=readme-ov-file#-api-endpoints)
- [Run Locally](https://github.com/nandarahmat/my-fiber-api?tab=readme-ov-file#-run-locally)
- [Testing with Postman](https://github.com/nandarahmat/my-fiber-api?tab=readme-ov-file#-testing-with-postman)

## ✨ Features

✅ **User Authentication (Register, Login)**  
✅ **Role-Based Access Control (User & Admin)**  
✅ **Category Management (CRUD Operations)**  
✅ **Province & City Data (Geolocation Support)**  
✅ **User Profile & Address Management**  
✅ **Store (Toko) Management**  
✅ **Product Management (CRUD Operations)**  
✅ **Transaction System**

## 🛠 Tech Stack

**Framework:** Fiber

**ORM:** GORM

**Database:** MYSQL

**Authentication:** JWT

**API Testing:** Postman

**Environment:** Go

## 🌐 API Endpoints

**Base URL:** `/api/v1`

---

### 🔐 Authentication

| Method | Endpoint         | Description       | Access |
| :----- | :--------------- | :---------------- | :----- |
| `POST` | `/auth/register` | Register new user | Public |
| `POST` | `/auth/login`    | User login        | Public |

---

### 📂 Categories (Admin Only for Write Operations)

| Method | Endpoint        | Description          | Access |
| :----- | :-------------- | :------------------- | :----- |
| `GET`  | `/category`     | List all categories  | Auth   |
| `GET`  | `/category/:id` | Get category details | Auth   |
| `POST` | `/category/`    | Create new category  | Admin  |
| `PUT`  | `/category/:id` | Update category      | Admin  |
| `DEL`  | `/category/:id` | Delete category      | Admin  |

---

### 🌏 Provinces & Cities

| Method | Endpoint                       | Description               | Access |
| :----- | :----------------------------- | :------------------------ | :----- |
| `GET`  | `/provcity/listprovincies`     | Get list of provinces     | Auth   |
| `GET`  | `/provcity/detailprovince/:id` | Get province details      | Auth   |
| `GET`  | `/provcity/listcities/:id`     | Get cities by province ID | Auth   |
| `GET`  | `/provcity/detailcity/:id`     | Get city details          | Auth   |

---

### 👤 User Management

| Method | Endpoint           | Description             | Access |
| :----- | :----------------- | :---------------------- | :----- |
| `GET`  | `/user`            | Get user data           | Auth   |
| `PUT`  | `/user`            | Update user             | Auth   |
| `GET`  | `/user/alamat`     | Get user addresses      | Auth   |
| `GET`  | `/user/alamat/:id` | Get user address by ID  | Auth   |
| `POST` | `/user/alamat`     | Create new user address | Auth   |
| `PUT`  | `/user/alamat/:id` | Update user address     | Auth   |
| `DEL`  | `/user/alamat/:id` | Delete user address     | Auth   |

---

### 🏪 Toko (Store) Management

| Method | Endpoint   | Description          | Access |
| :----- | :--------- | :------------------- | :----- |
| `GET`  | `/toko/my` | Get my store details | Auth   |
| `PUT`  | `/toko/my` | Update my store      | Auth   |
| `GET`  | `/toko`    | Get all stores       | Auth   |
| `GET`  | `/toko/id` | Get store by ID      | Auth   |

---

### 📦 Product Management

| Method | Endpoint       | Description        | Access |
| :----- | :------------- | :----------------- | :----- |
| `GET`  | `/product`     | Get all products   | Auth   |
| `GET`  | `/product/:id` | Get product by ID  | Auth   |
| `POST` | `/product/`    | Create new product | Admin  |
| `PUT`  | `/product/:id` | Update product     | Admin  |
| `DEL`  | `/product/:id` | Delete product     | Admin  |

---

### 💳 Transactions (Trx)

| Method | Endpoint   | Description            | Access |
| :----- | :--------- | :--------------------- | :----- |
| `GET`  | `/trx`     | Get all transactions   | Auth   |
| `GET`  | `/trx/:id` | Get transactions by ID | Auth   |
| `POST` | `/trx/`    | Create new transaction | Auth   |

---

#### 🔑 Keterangan Hak Akses:

- `Auth` → Hanya memerlukan login (bisa user biasa/admin).
- `Admin` → Hanya admin yang bisa akses.

## 🚀 Run Locally

Clone the project

```bash
  git clone https://github.com/nandarahmat/my-fiber-api.git
  cd my-fiber-api
```

Install dependencies

```bash
  go mod download
```

Configure `.env` file:

```bash
# Database
DB_HOST=localhost
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=myfiberdb
DB_PORT=3306

# JWT
JWT_SECRET=your_random_secret_key
```

Start MySQL Server, and start the server

```bash
  go run main.go
```

- **API Base URL:** `http://localhost:8000/api/v1`

## 🔍 Testing with Postman

1. Import the postman collection
2. Environment variables:

- `local` : `http://localhost:8000/api/v1`
- `token` : `(Login first to get token)`

**Sample Flow** :

1. `POST /auth/register` → Register test user
2. `POST /auth/login` → Auto set token into environment variables
