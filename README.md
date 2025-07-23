# xeodocs-dash-api
XeoDocs Project Dashboard API Service.

## How to Run

1.  **Ensure Go is installed**: Make sure you have Go installed on your system. You can download it from [golang.org](https://golang.org/doc/install).
2.  **Navigate to the project directory**: Open your terminal or command prompt and change to the project directory:

    ```bash
    cd /Users/fabian/Documents/CodeProjects/github.com/xeodocs/xeodocs-dash-api
    ```

3.  **Download dependencies**: Run `go mod tidy` to download all necessary dependencies:

    ```bash
    go mod tidy
    ```

4.  **Run the application**: Execute the `main.go` file:

    ```bash
    go run main.go
    ```

    The API service will start on `http://localhost:8080`.

## Endpoints

*   **GET /**: Returns a welcome message.
*   **GET /health**: Returns the health status of the service.
