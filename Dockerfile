FROM golang:latest
ENV GO111MODULE=on
WORKDIR /app/slack-push-notification
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD ["go","run","/app/slack-push-notification/main.go"]
