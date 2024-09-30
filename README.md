# Yuno-Card-Vault

## The Repository

API for managing card creation, using Go.

## Tech briefing

- Go: 1.20+
- MongoDB: latest

## Environments Availables

- [x] Develop

## Instructions

Test the API locally using VS Code.

All endpoints have token Beare auth, this token can only be used for 5 minutes.

### Local Development - Go

1. Navigate to the `yuno-cards` folder and run:

    ```bash
    go mod tidy
    ```
2. Navigate to the `yuno-cards` folder and create .env file, this file was sent by email.

3. Navigate to the `api-gateway` folder and run:

    ```bash
    go run main.go
    ```

4. Open other terminal and navigate the `auth` folder and run:

    ```bash
    go run main.go
    ```

5. Open other terminal and navigate the `cards` folder and run:

    ```bash
    go run main.go
    ```

6. Open your browser or Postman with this URI path: (postman collection send by email)

    ```
    http://localhost:8080/api/internal/healt
    ```


