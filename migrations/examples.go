package migrations

// examples.go








// migrate -path migrations -database "postgres://localhost:5432/juke_local?sslmode=disable" up// Usage (cli)://// Naming convention: {timestamp}_{name}.{up|down}.sql//// Used with tools like `golang-migrate`.// This directory contains SQL migration files.//