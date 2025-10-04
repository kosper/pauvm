FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .

# Build all executables
RUN go build -trimpath -ldflags="-s -w" -o /bin/ ./cmd/pauvm
RUN go build -trimpath -ldflags="-s -w" -o /bin/ ./cmd/pauven
RUN go build -trimpath -ldflags="-s -w" -o /bin/ ./cmd/paudiss

FROM alpine:3.19

COPY --from=builder /bin/pauvm /usr/local/bin/pauvm
COPY --from=builder /bin/pauven /usr/local/bin/pauven
COPY --from=builder /bin/paudiss /usr/local/bin/paudiss
