FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/volunteer-api ./cmd/api/main.go


FROM alpine:3.19

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /bin/volunteer-api .

EXPOSE 8080

CMD ["./volunteer-api"]