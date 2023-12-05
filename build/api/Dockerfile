FROM golang:latest

WORKDIR /app

# Copy necessary files
COPY ../../. .

RUN go mod download

# Install "air" for auto-reloading
RUN go install github.com/cosmtrek/air@latest

CMD go build -o ./tmp/cli ./cmd/cli && ./tmp/cli mup && air -c ./build/api/air.toml
