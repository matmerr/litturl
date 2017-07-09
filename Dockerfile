FROM golang:alpine
RUN apk add --update git nodejs
RUN go get github.com/matmerr/litturl
RUN cd $GOPATH/src/github.com/matmerr/litturl && \
	go build main.go && \
	cd client && \
        npm install && \
        npm run build
EXPOSE 8001
CMD ["sh", "-c", "litturl ${GOPATH}/src/github.com/matmerr/litturl/client/dist/"]
