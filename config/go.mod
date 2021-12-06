module github.com/config

go 1.17

replace github.com/utils => ../utils

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/utils v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.4.0 // indirect
