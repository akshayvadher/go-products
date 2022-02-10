# CRUD using
- Go
- Gin

# Products

# DB
- mariadb

## Prerequisite
- `docker-compose up`
- create `products` schema and create `products` table. Sql script provided in sql folder

# Run 
`go run .`

# Build and run prod
Windows
`go build .; $Env:GIN_MODE="release"; .\go-products.exe`
Linx
`go build .; export GIN_MODE="release"; .\go-products`
