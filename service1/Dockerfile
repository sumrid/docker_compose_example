# ============= basic ==============
FROM golang:1.12.5-alpine3.9
RUN apk update && apk add gcc git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build
EXPOSE 80

CMD ["./app"]


# =========     ========
# FROM alpine:3.10
# COPY app .
# EXPOSE 80
# CMD [ "./app" ]