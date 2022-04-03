FROM golang:1.17.5
RUN mkdir -p /opt/bombardier-go
WORKDIR /opt/bombardier-go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor
FROM alpine:latest
LABEL maintainer="scatolone" 
RUN apk add --no-cache curl
RUN apk add --no-cache tzdata
RUN apk --no-cache add ca-certificates
RUN cp /usr/share/zoneinfo/Europe/Rome /etc/localtime
RUN echo "Europe/Rome" >  /etc/timezone
RUN apk del --no-cache tzdata
WORKDIR /root/
COPY --from=0 /opt/bombardier-go/bombardier-go .
CMD ["./bombardier-go"]