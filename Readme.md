# CRUD using
- Go
- Gin

# Products

# DB
- mariadb

## Prerequisite
- `docker-compose up`
- It will automatically create required tables

# Run 
`go run .`

# Build and run prod
Windows
`go build .; $Env:GIN_MODE="release"; .\go-products.exe`
Linux
`go build .; export GIN_MODE="release"; .\go-products`
