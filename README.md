# Service Rest API Echo Golang

Ini adalah Service yang menggunakan [Echo Framework](https://echo.labstack.com/) di Go

## List of Contents
- [Jawaban Essay](#jawaban-essay)
- [Overview](#overview)
- [Request Validator](#request-validator)
- [JWT Token Middleware](#jwt-token-middleware)
- [Role Base Access Control](#role-base-access-control)
- [Object Relation Mapping](#object-relation-mapping)
- [Logging](#logging)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)

## Jawaban Essay
1. `Project Review` adalah proses untuk mengevaluasi tim terhadap kinerja tim pada project yang sedang dikerjakan ataupun yang sudah selesai, lalu `Project Planning` adalah tahap - tahap yang direncanakan untuk project yang akan dibuat ataupun yang sedang berjalan, biasa nya saya menggunakan tools seperti ClickUp, Trello ataupun Redmine untuk memonitor dan review kinerja tim.

2. Kebetulan saya belum berkesempatan mengelola vps AWS EC2, apabila pemahaman tentang `Load Balancer` adalah komponen untuk mendistribusikan lalu lintas jaringan ke service backend dan saya biasa nya menggunakan Nginx. Kemudian `Security Group` adalah firewall virtual untuk mengontrol lalu lintas masuk dan keluar, kalau pada vps Digital Ocean ada Dashboard tersendiri untuk mengaturnya, apabila langsung dari server bisa menggunakan UFW (Uncomplicated Firewall) untuk mengontrol port maupun IP Address.

3. Monitor penggunaan memory dengan runtime/pprof setelah di analisa dan mendapatkan fungsi atau library yang tidak direlease oleh memory maka langkah selanjutkan analisa kode tersebut apakah dari variable global, gourutine yang bocor ataupun tidak menutup fungsi defer close lalu perbaiki masalah nya dan monitor kembali bisa menggunakan htop ataupun runtime/pprof.

## Overview

Service API ini adalah contoh penerapan `Echo` di Go. Service ini mendukung autentikasi berbasis `JWT` dan `RBAC` terpusat untuk melindungi endpoint, validasi request yang kuat menggunakan `validator` dan menerapkan standard response success maupun response error. Service ini menggunakan Database `PostgreSQL` untuk penyimpanan data nya. Ketika service running maka akan `AutoMigrate` table `users`, `brands`, `products` kemudian pada table `users` otomatis seeding record untuk kebutuhan `Login` agar bisa mendapatkan akses pada endpoint / routes

## Request Validator

Service ini menggunakan [go-playground/validator](https://github.com/go-playground/validator) untuk validasi input data. Validator ini memastikan bahwa data yang masuk sesuai dengan aturan yang telah ditentukan, sehingga dapat mengurangi kesalahan dari data yang tidak valid.

- **Penggunaan `ctx.Validate`**: Semua request diproses dengan `ctx.Validate` untuk memeriksa data sebelum masuk ke layer service.

## JWT Token Middleware

Service ini menerapkan middleware `JWT Token` pada endpoint `Product` dan endpoint `Brand`.

## Role Base Access Control

Service ini menerapkan middleware `RBAC (Role Base Access Control)` untuk lebih jelas dan aman dalam mengakses endpoint `Product` maupun endpoint `Brand`.

## Object Relation Mapping

Service ini menggunakan [GORM](https://gorm.io/) sebagai Object-Relational Mapping (ORM) untuk mengelola interaksi dengan database PostgreSQL. GORM memudahkan penanganan operasi database seperti CRUD (Create, Read, Update, Delete) serta mendukung relasi antara model data, migrasi, dan query yang kompleks.

## Logging

Service ini menggunakan system `Logging` yang disimpan pada file secara periodik kemudian mencatat baik Request API maupun Response API untuk memudahkan analisa issue / bug.

### Requirements

- **Go 1.23 atau lebih baru**
- **PostgreSQL >= v16.14.2** untuk penyimpanan data

## Getting Started

Running aplikasi di local
Ikuti langkah-langkah berikut untuk memulai Service ini.
- **cd /path/to/project**
- **git clone https://github.com/robbyparlan/service-api.git**
- **cd /service-api**
- **cp .env.example .env**
- **cp config.yml config-local.yml** (Sesuaikan credential Database sesuai local anda pada file config-local.yml)
- **Pastikan Database PostgreSQL sudah running**
- **go run main.go**

## API Documentation

Service ini menggunakan `Postman Collection` untuk mendokumentasikan daftar API seperti `Login`, `Product`, `Brand`. Lalu untuk keterangan penggunaan, `sample request` dan kondisi `JWT` maupun `RBAC` sudah tertera pada file `Postman Collection`. Berikut file nya [Postman Collection](./SERVICE-API.postman_collection.json) kemudian import pada aplikasi [Postman](https://www.postman.com/downloads/).

