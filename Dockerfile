FROM node:18 as builder1

WORKDIR /build
COPY ./web .
RUN npm install && npm run build


FROM golang AS builder2
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build
ADD go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=builder1 /build/dist ./cmd/adapter-service/dist
RUN go build -ldflags "-s -w -extldflags '-static'" -o tiku-adapter ./cmd/adapter-service

FROM alpine

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder2 /build/tiku-adapter /app/tiku-adapter

# 创建数据目录用于持久化数据库文件
RUN mkdir -p /app/data

EXPOSE 8060
WORKDIR /app

# 声明数据卷用于持久化
VOLUME ["/app/data"]

ENTRYPOINT ["/app/tiku-adapter"]