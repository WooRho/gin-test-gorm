FROM golang:alpine as builder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
ENV GOOS linux
Run sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

WORKDIR /gin-test-gorm
COPY . .
RUN go build .

FROM scratch
WORKDIR /app
COPY --from=builder /gin-test-gorm/gin-test-gorm /app
COPY --from=builder /gin-test-gorm/config.yaml /app
EXPOSE 3003

CMD [ "./gin-test-gorm" ]