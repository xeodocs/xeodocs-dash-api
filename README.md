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

## Database Migrations

To add a new migration file using the Atlas tool, follow these steps:

1. **Ensure Atlas is installed**: Make sure you have the Atlas CLI installed. If not, you can install it by following the instructions on the [Atlas website](https://atlasgo.io).

2. **Generate a new migration file**: Use the Atlas CLI to create a new migration file. Replace `migration_name` with a descriptive name for your migration:
   ```bash
   atlas migrate new migration_name --dir "file://migrations"
   ```
   This will create a new file in the `migrations` directory with a timestamp and the name you provided.

3. **Edit the migration file**: Open the newly created migration file in your preferred editor and define the schema changes or SQL statements needed for this migration.

4. **Apply the migration**: The migrations are automatically applied in development when the container is started with `docker compose up --build -d`, and in production when you push to the main branch.

These steps will help you manage database schema changes effectively using Atlas.

## Endpoints

*   **GET /**: Returns a welcome message.
*   **GET /health**: Returns the health status of the service.
