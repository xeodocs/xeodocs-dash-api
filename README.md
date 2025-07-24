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

## API Endpoints

The API provides comprehensive CRUD operations for users, websites, and pages with session-based authentication.

### Authentication

- **POST /api/v1/auth/login**: Login with email and password
- **POST /api/v1/auth/logout**: Logout (requires authentication)
- **GET /api/v1/auth/me**: Get current user info (requires authentication)

### Users

All user endpoints require authentication except login.

- **GET /api/v1/users**: Get all users
- **GET /api/v1/users/:id**: Get user by ID
- **POST /api/v1/users**: Create new user
- **PUT /api/v1/users/:id**: Update user
- **DELETE /api/v1/users/:id**: Delete user

### Websites

All website endpoints require authentication.

- **GET /api/v1/websites**: Get all websites
- **GET /api/v1/websites/:id**: Get website by ID
- **GET /api/v1/websites/slug/:slug**: Get website by slug
- **POST /api/v1/websites**: Create new website
- **PUT /api/v1/websites/:id**: Update website
- **DELETE /api/v1/websites/:id**: Delete website

### Pages

All page endpoints require authentication.

- **GET /api/v1/pages**: Get all pages (supports ?website_id=X query parameter)
- **GET /api/v1/pages/:id**: Get page by ID
- **GET /api/v1/pages/slug/:slug**: Get page by slug
- **POST /api/v1/pages**: Create new page
- **PUT /api/v1/pages/:id**: Update page
- **DELETE /api/v1/pages/:id**: Delete page

### Health Check

- **GET /health**: Returns service health status

## Authentication

The API uses session-based authentication with Bearer tokens. Include the session token in the Authorization header:

```
Authorization: Bearer <session_token>
```

## Environment Configuration

- **Development**: Uses SQLite database at `../local/db.db`
- **Production**: Configured for Turso database (currently using SQLite as fallback)

### Environment Variables

- `ENVIRONMENT`: Set to "prod" for production mode (default: "dev")
- `TURSO_AUTH_TOKEN`: Authentication token for Turso database (production only)
- `PORT`: Server port (default: "8080")

## Project Structure

The project follows a clean architecture pattern:

```
├── api/                    # API layer (handlers, middleware, routes)
├── cmd/api/               # Application entry point
├── config/                # Configuration and database setup
├── internal/              # Private application code
│   ├── models/           # Data models and DTOs
│   ├── repository/       # Data access layer
│   └── service/          # Business logic layer
├── pkg/utils/            # Shared utilities
└── migrations/           # Database migration files
```
