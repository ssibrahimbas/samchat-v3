FROM golang:1.19-alpine
WORKDIR /
COPY ./services/auth ./auth
COPY ./services/shared.go ./shared.go
WORKDIR /auth
RUN go mod download && go mod tidy
RUN go build -o main ./src/main.go
CMD ["./main"]
