# Juke Backend

## Project Structure

This project follows a standard Clean Architecture layout for Go.

- **cmd/api**: Entry point of the application (`main.go`).
- **config**: Configuration loading logic.
- **internal**: Private application code.
  - **handler**: HTTP transport layer.
  - **service**: Business logic layer.
  - **repository**: Data access layer.
  - **model**: Domain models.
- **pkg**: Public library code.
  - **database**: Database connection logic.
- **migrations**: SQL migration files.

## Environment Setup

The project uses `.env` files for configuration.

- `.env.local`
- `.env.staging`
- `.env.prod`

## Running the Application

You can use the `Makefile` to run the application in different environments.

```bash
# Run in local mode (default)
make run

# Run in staging mode
make run ENV=staging
```

## Database Migrations

This project is set up to use `golang-migrate`.

1. Install `golang-migrate`:

   ```bash
   brew install golang-migrate
   ```
2. Run migrations:

   ```bash
   make migrate-up ENV=local
   ```
