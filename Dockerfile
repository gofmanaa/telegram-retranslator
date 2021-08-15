FROM alpine:3.7
COPY main main
ENTRYPOINT ["sh -h", "main"]