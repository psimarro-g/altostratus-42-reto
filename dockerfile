FROM golang:1.22.4 AS build

#download go modules
RUN mkdir /app
COPY go.* /app/

#set destination for COPY
WORKDIR /app
RUN go mod download

#copy source code
COPY . /app

#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/altostratus-42-reto-go -ldflags "-x main.build=." ./cmd

#run go binary in alpine
FROM alpine:3.18

#copy binary from build
COPY --from=build /app/bin/altostratus-42-reto-go /app/altostratus-42-reto-go
WORKDIR /app

RUN chmod +x altostratus-42-reto-go

EXPOSE 8080

CMD "./altostratus-42-reto-go"