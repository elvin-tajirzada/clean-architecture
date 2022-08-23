FROM golang:1.18.3-alpine3.16
WORKDIR /clean-architecture
COPY . /clean-architecture
RUN go build -o main ./cmd/clean-architecture
EXPOSE 8008
CMD ["/clean-architecture/main"]