## First stage build
FROM golang:1.16.5 as build

RUN mkdir -p /opt/app
WORKDIR /opt/app

RUN go get -u github.com/swaggo/swag/cmd/swag
COPY . .
RUN swag init -g cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux  go build -o template-project-ms -mod=vendor cmd/main.go

## Second stage build
FROM alpine:latest
LABEL maintainer="TDGroup Italia Srl" 
RUN apk add --no-cache curl
RUN apk add --no-cache tzdata
RUN apk --no-cache add ca-certificates

RUN cp /usr/share/zoneinfo/Europe/Rome /etc/localtime
RUN echo "Europe/Rome" >  /etc/timezone
##RUN apk del --no-cache tzdata

RUN adduser -H -D -s /sbin/nologin cochise
RUN adduser cochise cochise

WORKDIR /opt/app
COPY --from=build /opt/app/template-project-ms .
RUN mkdir -p docs/swagger/
COPY --from=build /opt/app/docs/swagger.json ./docs/swagger/
COPY --from=build /opt/app/docs/version.json ./docs/
RUN chown -R cochise.cochise /opt/app/
CMD ["./template-project-ms"]
HEALTHCHECK --timeout=5s --start-period=10s --interval=60s CMD curl http://localhost:9091/${HEALTH_BASE_PATH:-template-project-ms}/health || exit 1
