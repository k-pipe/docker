#FROM golang:1.19
FROM golang:1.22
#as builder

WORKDIR /app

COPY go.mod main.go ./

#RUN GOOS=linux GOARCH=amd64 go build \
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN chmod u+x /app/main
#FROM scratch

#COPY --from=builder /app/main /main
#RUN chmod u+x /main

ENTRYPOINT ["/app/main"]