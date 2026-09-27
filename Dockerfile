FROM golang:1.27 AS builder

WORKDIR /app/backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

RUN CGO_ENABLED=0 go build -o /app/one-small-step .


FROM debian:13-slim

WORKDIR /app

COPY --from=builder /app/one-small-step ./one-small-step
COPY frontend ./frontend
COPY config ./config

ENV CONFIG_PATH=/app/config/config.yaml
ENV FRONTEND_PATH=/app/frontend

EXPOSE 8080

CMD ["./one-small-step"]
