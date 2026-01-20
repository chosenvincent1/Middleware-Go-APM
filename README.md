# Go Middleware Demo
This is a simple Go application demonstrating **error handling and observability** using [Middleware](https://middleware.io) APM. It includes automatic error tracking, distributed tracing, and request context capture for HTTP handlers.

The app simulates a small user service with:

- `GET /user?id=<id>` – Fetch a user by ID
- `POST /user/create` – Create a new user

Middleware automatically captures errors, stack traces, request metadata, and can also attach custom data for observability.

---

## Getting Started

### Prerequisites

- Go 1.17+
- A [Middleware](https://middleware.io) account and API credentials (access token & target)

### Installation

Clone the repository:

```bash
git clone https://github.com/<your-username>/middleware-go-demo.git
cd middleware-go-demo
```

#### Install dependencies

```bash
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-http/http
```

### Running Your Application

1. Replace `<your-access-token>` and `<your-target-url>` in `main.go` with your Middleware credentials.
2. Start the server:
```bash
go run main.go
```
3. The server runs on `http://localhost:8080`

### Testing Error Tracking

Fetch a user that doesn’t exist:

```bash
curl http://localhost:8080/user?id=999
```

This will return a 404 and Middleware will automatically capture the error with full context in your dashboard.