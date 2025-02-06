# install the dependencies
FROM golang:1.23.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

# copy all the project
COPY . .

# build the app
RUN go build -o main .

# runn the app
FROM alpine:latest AS runner

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 8080

ENV MONGODB_URI=${MONGODB_URI}
ENV SERVER_LOG=${SERVER_LOG}
ENV SERVER_PORT=${SERVER_PORT}



CMD ["./main"]