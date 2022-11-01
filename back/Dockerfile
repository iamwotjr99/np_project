FROM golang:alpine AS builder

WORKDIR /back

COPY go.mod go.sum server.go ./

RUN go mod download

COPY . ./

RUN ls

RUN go build -o /webServer

RUN rm -rf /back/server.go

################################

FROM alpine

EXPOSE 8080

WORKDIR /back

COPY --from=builder /webServer /backend

ENTRYPOINT ["/backend"]