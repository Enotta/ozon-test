FROM golang:latest

# Set destination for COPY/RUN
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /coffee server.go

RUN go get github.com/graphql-go/graphql
RUN go get github.com/lib/pq

# Expose port
EXPOSE 8080

# Run
CMD ["/coffee"]