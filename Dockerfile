FROM golang:1.17-alpine as build
COPY . src
WORKDIR src
RUN make buil


FROM alpine:3.7
COPY --from=build /main /var/main
COPY /.env /var/.env

ENTRYPOINT ["sh -h", "/var/main"]