# Build the app binary
FROM golang:1.20 as builder

WORKDIR /app

# Copy and install go module dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server ./cmd/...

FROM gcr.io/distroless/static:debug

WORKDIR /app
COPY --from=builder /app/server .
COPY config.yaml .

ENTRYPOINT [ "./server" ]
CMD [ "-config", "config.yaml" ]
