FROM golang:1.22.3 AS builder
WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG PROGRAM=executable
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build

FROM alpine:3.10.3 AS runner
WORKDIR /app
RUN apk update && apk add --no-cache tzdata
ARG PROGRAM=executable
COPY --from=builder /builder/${PROGRAM} .
ENV PROGRAM=${PROGRAM}
CMD ./$PROGRAM
