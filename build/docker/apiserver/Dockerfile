# docker build -t=jdxj/apiserver:v0.3.0 .
# docker push jdxj/apiserver:v0.3.0

FROM golang:alpine as build
RUN apk --no-cache add tzdata

FROM alpine
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
RUN mkdir -p /sign
COPY apiserver.run /sign/apiserver.run
