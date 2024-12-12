# Rate Limiter Challenge

### Objective

Develop a rate limiter in Go that restricts requests based on IP address or API token, configurable via .env file.

### Key Features

- IP-Based Limiting: Restrict requests per IP within a time window.
- Token-Based Limiting: Restrict requests per unique token (API_KEY header). Token-specific limits override IP limits.
- Blocking Time: Configurable block duration when limits are exceeded.
- Redis Integration: Rate limiter logic is backed by Redis for fast access and scalability.
- Middleware: Designed as middleware to integrate seamlessly into web servers.
- Custom Strategies: Allows swapping Redis with other storage backends easily.

- **How to Run Locally:**

### Using Go

1. Clone the repository.
2. Ensure Redis is running on localhost:6379.
3. Run

```bash
go run cmd/main.go
```

### Using Docker-Compose

1. Clone the repository.
2. Run

```bash
docker-compose up --build
```

### Configuration

Update configs/config.env with:

- IP and token limits.
- Block durations for IPs and tokens

### Testing the Rate Limiter

- Routes
  /token: Requires a valid API_KEY header; limits requests by token.
  /ip: No API_KEY required; limits requests by IP.
  /both:
  If no API_KEY is sent, uses IP-based limits.
  If API_KEY is sent, validates using token-based limits.

- Logs
  Logs show request origins (IP or token) and their access count.
  Logs indicate if access is blocked.

### Running Tests

1. Ensure Redis is running on localhost:6379.
2. Navigate to tests and execute:

```bash
go test -v
```
