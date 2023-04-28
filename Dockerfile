FROM golang:alpine
ENV PORT=8080
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE $PORT
