# MICROSERVICE-ECOMMERCE API
proyek ini merupakan backend e-commerce yang dibangun dengan arsitektur **microservices** menggunakan **Golang**, **gRPC**, dan **REST API**. 
Sistem dirancang modular dengan service terpisah seperti `user-service`, `product-service`, `cart-service`. untuk authentication saya menggunakan yang simple berupa session based auth yang disimpan sebagai http only cookies.


## How to run this Project

- Install Docker Desktop terlebih dahulu jika belum tersedia di perangkatmu → https://www.docker.com/products/docker-desktop

- Jalankan Docker Desktop hingga status-nya "Running".

- Buat file .env di root folder proyek (di sebelah docker-compose.yml), lalu isi dengan konfigurasi contoh di paling bawah readme ini.

- Jalankan perintah berikut untuk membangun dan menjalankan semua service (app, MySQL, Redis, dsb):

```
docker-compose up -d --build

```

- Tunggu sampai container berhasil berjalan. Jika berhasil, akses API di port local yang telah didefinisikan pada .env

- Gunakan Postman atau Curl untuk menguji endpoint-endpoint API.

### Catatan :

- file seeders akan otomatis dibuat pada saat dijalankan bantuan untuk pengujian API sudah terdapat didalam root project bernama api-test.json
- file untuk pembuatan schema database permasing-masing service sudah ada di init.sql

## Route yang Tersedia dan Penjelasannya

Berikut adalah daftar lengkap endpoint API yang tersedia dalam proyek ini, dikelompokkan berdasarkan fungsionalitas dan otorisasi akses.

# 👤 User Service API

Dokumentasi endpoint untuk service user yang mencakup fitur autentikasi, profil, dan manajemen alamat.

---
# User-services

## Auth Routes

**Base URL:** `/api/auth`

| Method | Endpoint                         | Akses     | Deskripsi                                  |
|--------|----------------------------------|-----------|--------------------------------------------|
| POST   | `/api/auth/login`               | Public    | Login user                                  |
| POST   | `/api/auth/logout`              | Public    | Logout user                                 |
| POST   | `/api/auth/register`            | Public    | Registrasi user baru                        |
| POST   | `/api/auth/reset-password`      | Public    | Reset password (dengan token)               |
| POST   | `/api/auth/forgot-password`     | Public    | Minta reset password (via email)            |
| GET    | `/api/auth/me`                  | Private   | Mendapatkan info user yang login            |
| PUT    | `/api/auth/change-password`     | Private   | Ganti password                              |

### Admin Auth Routes

**Base URL:** `/api/auth/admin`

| Method | Endpoint                     | Akses     | Deskripsi                                  |
|--------|------------------------------|-----------|--------------------------------------------|
| GET    | `/api/auth/admin/user`       | Admin     | Mendapatkan semua data user                |
| GET    | `/api/auth/admin/user/:id`   | Admin     | Mendapatkan data user berdasarkan ID       |

---

## User Routes

**Base URL:** `/api/user`  
> Semua route ini memerlukan autentikasi token (`AuthRequired()`)

###  Profile

| Method | Endpoint              | Akses   | Deskripsi                         |
|--------|-----------------------|---------|-----------------------------------|
| GET    | `/api/user/profile`   | Private | Mendapatkan data profil user      |
| PUT    | `/api/user/profile`   | Private | Memperbarui data profil user      |

### Address

| Method | Endpoint                                | Akses   | Deskripsi                                |
|--------|-----------------------------------------|---------|------------------------------------------|
| GET    | `/api/user/addresses`                   | Private | Mendapatkan semua alamat user            |
| POST   | `/api/user/addresses`                   | Private | Menambahkan alamat baru                  |
| PUT    | `/api/user/addresses/:id`               | Private | Memperbarui alamat berdasarkan ID        |
| DELETE | `/api/user/addresses/:id`               | Private | Menghapus alamat berdasarkan ID          |
| PUT    | `/api/user/addresses/:id/set-main`      | Private | Menjadikan alamat tersebut sebagai utama |

---
# Product-services
## Product Routes

**Base URL:** `/api/products`

| Method | Endpoint                             | Akses  | Deskripsi                                     |
|--------|--------------------------------------|--------|-----------------------------------------------|
| GET    | `/api/products`                      | Public | Mendapatkan semua produk                      |
| GET    | `/api/products/:slug`                | Public | Mendapatkan detail produk berdasarkan slug    |
| POST   | `/api/products/admin`                | Admin  | Membuat produk baru                           |
| PUT    | `/api/products/admin/:id`            | Admin  | Memperbarui produk                            |
| DELETE | `/api/products/admin/:id`            | Admin  | Menghapus produk                              |
| POST   | `/api/products/admin/upload-local`   | Admin  | Memperbarui produk                            |
| GET    | `/api/products/admin/:id/download`   | Admin  | Menghapus produk                              |


---

##  Category Routes

**Public Endpoint:** `/api/categories`  
**Admin Endpoint:** `/api/admin/categories`

| Method | Endpoint                                          | Akses  | Deskripsi                               |
|--------|---------------------------------------------------|--------|-----------------------------------------|
| GET    | `/api/categories`                                 | Public | Mendapatkan semua kategori              |
| POST   | `/api/admin/categories`                           | Admin  | Membuat kategori baru                   |
| PUT    | `/api/admin/categories/:id`                       | Admin  | Memperbarui kategori                    |
| DELETE | `/api/admin/categories/:id`                       | Admin  | Menghapus kategori                      |
| POST   | `/api/admin/categories/:id/subcategories`         | Admin  | Menambah subkategori ke kategori        |
| PUT    | `/api/admin/categories/subcategories/:subId`      | Admin  | Memperbarui subkategori                 |
| DELETE | `/api/admin/categories/subcategories/:subId`      | Admin  | Menghapus subkategori                   |

---

##  Variant Routes

**Base URL:** `/api/variants`

| Method | Endpoint                                | Akses  | Deskripsi                                 |
|--------|-----------------------------------------|--------|-------------------------------------------|
| GET    | `/api/variants`                         | Public | Mendapatkan semua variant types           |
| POST   | `/api/variants`                         | Admin  | Membuat variant type                      |
| PUT    | `/api/variants/:id`                     | Admin  | Memperbarui variant type                  |
| DELETE | `/api/variants/:id`                     | Admin  | Menghapus variant type                    |
| POST   | `/api/variants/:id/values`              | Admin  | Menambahkan value ke variant              |
| PUT    | `/api/variants/values/:valueId`         | Admin  | Memperbarui value dari variant            |
| DELETE | `/api/variants/values/:valueId`         | Admin  | Menghapus value dari variant              |
| POST   | `/api/variants/map/category`            | Admin  | Mapping variant ke kategori               |
| POST   | `/api/variants/map/subcategory`         | Admin  | Mapping variant ke subkategori            |

---

##  Attribute Routes

**Base URL:** `/api/attributes`

| Method | Endpoint                                 | Akses  | Deskripsi                                 |
|--------|------------------------------------------|--------|-------------------------------------------|
| GET    | `/api/attributes`                        | Public | Mendapatkan semua attribute               |
| POST   | `/api/attributes`                        | Admin  | Membuat attribute baru                    |
| PUT    | `/api/attributes/:id`                    | Admin  | Memperbarui attribute                     |
| DELETE | `/api/attributes/:id`                    | Admin  | Menghapus attribute                       |
| POST   | `/api/attributes/:id/values`             | Admin  | Menambahkan value ke attribute            |
| PUT    | `/api/attributes/values/:valueId`        | Admin  | Memperbarui value dari attribute          |
| DELETE | `/api/attributes/values/:valueId`        | Admin  | Menghapus value dari attribute            |

---

#  Cart Service API

---
## Cart Routes
**Base URL:** `/api/cart`

| Method | Endpoint              | Akses   | Deskripsi                                  |
|--------|-----------------------|---------|--------------------------------------------|
| GET    | `/api/cart`           | Private | Mendapatkan isi keranjang user             |
| POST   | `/api/cart`           | Private | Menambahkan item ke keranjang              |
| PUT    | `/api/cart/items/:id` | Private | Memperbarui jumlah item dalam keranjang    |
| DELETE | `/api/cart/items/:id` | Private | Menghapus item dari keranjang              |
| DELETE | `/api/cart`           | Private | Menghapus semua item dari keranjang        |

---


### Catatan :
**pastikan semua configurasi .env yang ada di setiap folder service sudah disetting sesuai .env.example**
