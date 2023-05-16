# Build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
RUN chmod +x start.sh
RUN chmod +x wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]