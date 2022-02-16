# What
It is just a POC to verify following 
- HTTP Api works
- DB integration works
- CRUD end to end
- Pagination works
- Custom errors can be sent
- DB migration works
- Configuration file can be setup
- Configuration can be overridden from ENV variables

Basically, in order to create a PROD grade app, all above things are necessary

# Tech
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
