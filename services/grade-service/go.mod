module github.com/Ethea2/Distributed-Fault-Tolerance/services/grade-service

go 1.23.4

require (
	github.com/Ethea2/Distributed-Fault-Tolerance/services/common v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.2.1
	github.com/jackc/pgx/v5 v5.7.4
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/Ethea2/Distributed-Fault-Tolerance/services/common => ../common
