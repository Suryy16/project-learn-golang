# Project Learn Golang

Repository ini berisi berbagai project backend menggunakan Golang sebagai bagian dari perjalanan belajarku dalam menguasai pengembangan backend dengan Go.

## 📚 Tentang Repository

Repository ini dibuat untuk mendokumentasikan progres pembelajaran saya dalam pengembangan aplikasi backend menggunakan bahasa pemrograman Go (Golang). Setiap project di repository ini merupakan implementasi konsep-konsep yang berbeda dalam ekosistem backend development.

## 🎯 Tujuan Pembelajaran

- Memahami fundamental bahasa pemrograman Go
- Menguasai konsep-konsep backend development
- Belajar best practices dalam penulisan kode Go
- Mengimplementasikan berbagai design patterns
- Membangun aplikasi backend yang scalable dan maintainable

## 📂 Daftar Project

### 1. Task CLI (`task-cli/`)

Aplikasi command-line untuk manajemen task/todo list yang dibangun dengan Go. Project ini mencakup:

- **Fitur Utama:**
  - Menambah task baru
  - Menampilkan daftar task
  - Mengupdate status task
  - Menghapus task
  - Persistent storage menggunakan JSON

- **Konsep yang Dipelajari:**
  - Command-line flag parsing
  - File I/O operations
  - JSON serialization/deserialization
  - Struct dan methods di Go
  - Error handling
  - Unit testing

- **Dependencies:**
  - `github.com/aquasecurity/table` - untuk menampilkan data dalam format tabel

**Cara Menjalankan:**
```bash
cd task-cli
go run main.go -add "Nama task baru"
go run main.go -list
```

## 🚀 Instalasi

### Prerequisites

- Go 1.24.6 atau lebih tinggi
- Git

### Setup

1. Clone repository ini:
```bash
git clone <repository-url>
cd project-learn-golang
```

2. Masuk ke direktori project yang ingin dijalankan:
```bash
cd task-cli
```

3. Install dependencies:
```bash
go mod download
```

4. Jalankan project:
```bash
go run main.go
```

## 📖 Struktur Project

Setiap project dalam repository ini memiliki struktur yang terorganisir:
- `main.go` - Entry point aplikasi
- `*.go` - File-file source code
- `*_test.go` - File-file unit test
- `go.mod` - Module dependencies
- `README.md` - Dokumentasi spesifik project (jika ada)

## 🧪 Testing

Untuk menjalankan test pada setiap project:

```bash
cd <nama-project>
go test -v
```

Untuk melihat test coverage:

```bash
go test -cover
```

## 📝 Catatan Pembelajaran

Repository ini akan terus diupdate dengan project-project baru seiring dengan progres pembelajaran. Setiap project dirancang untuk mengeksplorasi aspek yang berbeda dari pengembangan backend dengan Go.

## 🔜 Project Mendatang

Beberapa project yang akan ditambahkan:
- REST API dengan Gin/Echo framework
- Database integration (PostgreSQL/MongoDB)
- Authentication & Authorization
- Microservices architecture
- Message queue (RabbitMQ/Kafka)
- Containerization dengan Docker
- Dan masih banyak lagi...

## 📚 Sumber Belajar

- [Go Official Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

## 🤝 Kontribusi

Ini adalah repository pembelajaran pribadi, namun feedback dan saran selalu diterima dengan baik!

## 📄 Lisensi

Repository ini dibuat untuk tujuan pembelajaran.

---

⭐ Happy Learning! ⭐
