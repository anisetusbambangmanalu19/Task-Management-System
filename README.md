# Task Management System

Aplikasi manajemen proyek dan task berbasis web yang dibuat menggunakan Golang (Gin) sebagai backend dan React (Vite) sebagai frontend. Sistem ini mendukung autentikasi menggunakan JWT serta relasi database PostgreSQL.

## Fitur

* Register dan Login user
* Autentikasi menggunakan JWT
* CRUD Project
* CRUD Task dalam Project
* Validasi kepemilikan data (user hanya bisa mengakses project miliknya)
* Relasi database (User → Project → Task)

## Teknologi yang Digunakan

### Backend

* Golang
* Gin Framework
* PostgreSQL
* JWT Authentication
* Clean Architecture

### Frontend

* React (Vite)
* Axios
* React Router
* Tailwind CSS

### Database

* PostgreSQL dengan relasi foreign key

## Struktur Sistem

Relasi database:

User memiliki banyak Project
Project memiliki banyak Task

Struktur backend menggunakan pendekatan Clean Architecture:

* entity: struktur data/model
* repository: akses database
* usecase: business logic
* delivery: handler HTTP
* config: konfigurasi aplikasi

## Cara Menjalankan Backend

1. Pastikan PostgreSQL sudah berjalan
2. Buat database
3. Atur konfigurasi database di file environment
4. Jalankan:

```
go run cmd/main.go
```

Server akan berjalan di port yang telah dikonfigurasi.

## Cara Menjalankan Frontend

1. Masuk ke folder frontend
2. Install dependency:

```
npm install
```

3. Jalankan:

```
npm run dev
```

Frontend akan berjalan di [http://localhost:5173](http://localhost:5173) (default Vite).

## Alur Sistem

1. User melakukan register
2. User login dan mendapatkan token JWT
3. Token dikirim pada setiap request ke backend
4. Backend memvalidasi token
5. User hanya dapat mengakses dan mengelola project miliknya sendiri

## Tujuan Proyek

Proyek ini dibuat untuk mempelajari:

* Implementasi REST API dengan Golang
* Clean Architecture
* JWT Authentication
* Relasi database dan foreign key
* Integrasi backend dan frontend

---

