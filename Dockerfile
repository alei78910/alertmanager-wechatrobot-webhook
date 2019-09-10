FROM golang:latest
WORKDIR /app
COPY . .
RUN go get -u github.com/gin-gonic/gin
RUN go build -o main .
EXPOSE 8999
CMD ["./main"]