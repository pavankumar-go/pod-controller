FROM golang:1.13.0-alpine AS builder

WORKDIR /app
COPY . .

RUN  apk update \
    && apk add git \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-extldflags '-static'" -mod=vendor

FROM alpine
COPY --from=builder /app/pod-controller  /app/pod-controller
WORKDIR /app

ENTRYPOINT [ "./pod-controller" ]