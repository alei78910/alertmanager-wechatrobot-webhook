FROM golang:alpine
WORKDIR /app
COPY . .
RUN go get -u github.com/gin-gonic/gin
RUN go build -o main .
EXPOSE 8999
CMD ["./main"]

# FROM scratch
# ADD main /
# EXPOSE 8999
# CMD ["/main"]
# CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main .