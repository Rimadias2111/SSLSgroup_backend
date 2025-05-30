FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/myapp .

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/data /app/data

COPY --from=builder /app/myapp /myapp

EXPOSE 8080

CMD ["/myapp"]
