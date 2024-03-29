# First Level

FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -build="go build ." -log-prefix="false" -command="./altschool-sms"

# Second Level

#FROM alpine
#
#COPY --from=builder app/altschool-sms .
#COPY --from=builder app/log.json .
#COPY --from=builder app/config.env .
#
#CMD ["./altschool-sms"]