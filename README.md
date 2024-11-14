# Backend Project-based Test

## Deskripsi

Anda diminta untuk mengembangkan backend untuk sistem manajemen inventaris sebuah toko baju. Sistem ini harus dapat menangani pembaruan stok baju, penambahan baju baru, dan pencarian baju berdasarkan warna dan ukuran.

## Spesifikasi

- Sistem harus dapat menangani operasi CRUD untuk baju.
- Setiap baju memiliki atribut warna, ukuran, harga, dan stok.
- Sistem harus dapat mencari baju berdasarkan warna dan ukuran.
- Sistem harus dapat menambahkan stok baju.
- Sistem harus dapat mengurangi stok baju.
- Sistem harus dapat menampilkan semua baju yang tersedia.

### Bonus

- Sistem dapat menampilkan semua baju yang stoknya habis.
- Sistem dapat menampilkan semua baju yang stoknya kurang dari 5.

## Tech Stack

Dibebaskan untuk menggunakan tech stack apapun yang menurut Anda cocok untuk menyelesaikan tugas ini. Recommended stack: Node.js, Express.js, MongoDB or Go, Gin, Gorm, PostgreSQL.

## Deliverables

Silakan fork repository ini dan submit link repository hasil pengerjaan Anda ke [https://forms.gle/kAsjU8hFHyVNED2u6](https://forms.gle/kAsjU8hFHyVNED2u6).


# How to Run, and Needed

## Getting started

### Requirements
- go version >= 1.20
- Makefile
- mockery v2.32.0
- swag


## Usage

### Config
Clone config file `config.yaml.example` from directory `/config/files`, put it on the same directory and rename it to `config.yaml`

You can also define config from env.


### API Server
running http server using
```
make run-api
```
or go command
```
go run main.go serve-http
```



### DB Migration
Reference:
- [go-migrate](https://github.com/golang-migrate/migrate)
- [go-migrate for PostgreSQL](https://github.com/golang-migrate/migrate/tree/master/database/postgres)

#### Install Go Migrate
```
$ go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
#### Create Migration File
```
$ migrate create -ext sql -dir db/migrations [name_of_migration_file]
```
#### Migrate up
```
$ migrate -database "postgres://[user]:[password]@[host]:[port]/[dbname]?query" -path db/migrations up
```
#### Migrate down
```
$ migrate -database "postgres://[user]:[password]@[host]:[port]/[dbname]?query" -path db/migrations down
```


accessing swagger docs using
```
http://localhost:8900/docs/index.html
```
