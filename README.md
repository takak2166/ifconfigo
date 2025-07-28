# ifconfigo

A simple HTTP server that returns detailed IP address information.

## Features

- Returns X-Forwarded-For, X-Real-IP, and RemoteAddr separately
- Shows "not present" when headers are missing
- Supports proxy headers (X-Forwarded-For, X-Real-IP)
- Configurable port via command line argument
- Access logging

## Usage

### Basic usage
```bash
go run main.go
```

### Custom port
```bash
go run main.go 3000
```

### Build and run
```bash
go build -o ifconfigo
./ifconfigo
```

### Using Docker

#### Build and run with Docker
```bash
# Build the image
docker build -t ifconfigo .

# Run the container
docker run -p 8080:8080 ifconfigo
```

#### Using Docker Compose
```bash
# Start the service
docker-compose up -d

# Stop the service
docker-compose down
```

#### Custom port with Docker
```bash
# Run on custom port (e.g., 3000)
docker run -p 3000:8080 ifconfigo ./ifconfigo 3000
```

## Testing

Once the server is running, you can test it with:

```bash
# Direct connection
curl http://localhost:8080/

# With X-Forwarded-For header
curl -H "X-Forwarded-For: 192.168.1.100, 10.0.0.1" http://localhost:8080/

# With X-Real-IP header
curl -H "X-Real-IP: 203.0.113.1" http://localhost:8080/

# Through SOCKS proxy
curl --socks5 localhost:1080 http://localhost:8080/
```

## Output Format

The server returns the following information:

```
X-Forwarded-For: [value or "not present"]
X-Real-IP: [value or "not present"]
RemoteAddr: [IP address]
```

## How it works

The server extracts and displays all available IP-related information:
- **X-Forwarded-For**: Shows the original client IP when request comes through a proxy
- **X-Real-IP**: Shows the real client IP (often set by reverse proxies)
- **RemoteAddr**: Shows the direct connection IP address

This provides complete visibility into the IP address chain, useful for debugging proxy configurations and understanding request routing.