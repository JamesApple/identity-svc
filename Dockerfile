FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy && go mod download
COPY . .
RUN cd ./cmd/serve && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /app/cmd/serve/main .
CMD ["./main"] 


