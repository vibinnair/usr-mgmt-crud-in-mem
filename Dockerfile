FROM golang:1.22-alpine

#set working directory
WORKDIR /app/user-management-in-memory

# download dependencies
COPY go.mod .
RUN go mod download

# copy all go lang files
COPY cmd ./cmd
COPY internal ./internal
COPY model ./model

#build go lang
RUN go build -o user-mgmt-service cmd/main.go

EXPOSE 8282

CMD ["./user-mgmt-service"]