FROM golang:alpine as builder

ENV GOPROXY https://goproxy.cn/

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o goadmin .

FROM alpine as runner
COPY --from=builder /usr/src/app/goadmin /
COPY config /config
COPY temp /temp

EXPOSE 8000

CMD ["/goadmin","server","-c", "/config/settings.yml"]

