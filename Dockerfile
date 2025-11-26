FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o capital-gains ./cmd/cli

FROM gcr.io/distroless/static
COPY --from=builder /app/capital-gains /capital-gains
ENTRYPOINT ["/capital-gains"]
