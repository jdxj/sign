# docker build -t=jdxj/sign:v0.2.2 .
# docker push jdxj/sign:v0.2.2

FROM golang:alpine as build
RUN apk --no-cache add tzdata

FROM alpine
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
RUN mkdir -p /sign
COPY apiserver.out /sign/apiserver.out