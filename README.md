# simple-project-be-syarifuddin
Build REST API for product entity using Go and PostgreSQL

## Setup Database
- Install [PostgreSQL](https://www.postgresql.org/download/)
- Make database
- Import ```database/dagangan.pgsql``` ke database
- Konfigurasi ```database.go``` sesuai database yang dimiliki

## Setup Project
- Install [Go](https://go.dev/doc/install)
- Install dependencies
```
go get -u github.com/gorilla/mux
go get -u github.com/lib/pq
go get -u github.com/cosmtrek/air
```

## Compiles and hot-reloads for development
```
air
```

## REST API List
- `GET  /products?page=1&size=5&orderby=id&desc=true&hargaMin=10000&hargaMax=20000`
- `GET  /products/[jenis]?page=1&size=5&orderby=id&desc=true&hargaMin=10000&hargaMax=20000`
- `GET  /product/[id]`
- `POST /addProduct`
- `PUT  /updateProduct`
- `DEL  /deleteProduct`