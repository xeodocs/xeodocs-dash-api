# xeodocs-dash-api
XeoDocs Project Dashboard API Service.

## How to Run (Development)

This project uses Docker Compose for a consistent development environment.

1.  **Ensure Docker is installed**: Make sure you have Docker and Docker Compose installed on your system. You can download it from [docker.com](https://www.docker.com/get-started).

2.  **Build and start the development environment**: From the project root, run:

    ```bash
    docker compose up --build
    ```

    This will build the `Dockerfile.development` image, install dependencies, and start the Go application. The API service will be accessible at `http://localhost:8080`.

3.  **Access the container (optional)**: To run commands inside the development container (e.g., for migrations), open a new terminal and run:

    ```bash
    docker compose exec app bash
    ```

    From within the container, you can run Go commands, migration commands, etc.

## Endpoints

*   **GET /**: Returns a welcome message.
*   **GET /health**: Returns the health status of the service.

## Database Migrations (using golang-migrate/migrate)

This project uses `golang-migrate/migrate` for database migrations. The migration files are located in the `migrations` directory.

### Installation

To install the `migrate` CLI tool:

```bash
go install -tags 'sqlite3 libsql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Usage

**To create a new migration:**

```bash
migrate create -ext sql -dir migrations -seq <migration_name>
```

This will create two new files in the `migrations` directory: `<timestamp>_<migration_name>.up.sql` and `<timestamp>_<migration_name>.down.sql`. You should add your SQL schema changes to the `up.sql` file and the corresponding rollback SQL to the `down.sql` file.

**To apply migrations (migrate up):**

```bash
migrate -path migrations -database "sqlite://your_database.db" up
```

Replace `sqlite://your_database.db` with your actual database connection string (e.g., for Turso, it would be `libsql://your-db-name.turso.io?authToken=your-token`).

**To revert the last migration (migrate down):**

```bash
migrate -path migrations -database "sqlite://your_database.db" down 1
```

**To force a migration version (use with caution):**

```bash
migrate -path migrations -database "sqlite://your_database.db" force <version>
```
