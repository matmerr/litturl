FROM golang:alpine
RUN apk add --update git nodejs
RUN go get github.com/matmerr/litturl && \
        cd /go/src/github.com/matmerr/litturl && \
        go build main.go && \
        cd client && \
        npm install && \
        npm run build
EXPOSE 8001
CMD ["litturl"]