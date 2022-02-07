FROM golang as builder
WORKDIR /source
COPY . ./ 
RUN GOPROXY=https://goproxy.io  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o start

FROM alpine:latest
ENV RELEASE=true
ENV GIN_MODE=release
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /source
RUN echo "Asia/Shanghai" > /etc/timezone
COPY ./configs/config.yaml ./configs/config.yaml
COPY --from=builder /source/start ./
CMD ./start
