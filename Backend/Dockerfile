FROM golang
WORKDIR /var/app
COPY App/* .
RUN go build -o main

FROM golang
WORKDIR /var/app
COPY --from=0 /var/app .
CMD ["/var/app/main"]
