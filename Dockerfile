FROM golang:alpine
RUN apk add --update git nodejs
<<<<<<< HEAD
RUN go get github.com/matmerr/litturl && \
        cd /go/src/github.com/matmerr/litturl && \
        go build main.go && \
        cd client && \
        npm install && \
        npm run build
EXPOSE 8001
CMD ["litturl"]
=======
RUN go get github.com/matmerr/litturl
RUN cd $GOPATH/src/github.com/matmerr/litturl && \
	go build main.go && \
	cd client && \
        npm install && \
        npm run build
EXPOSE 8001
CMD ["litturl", "$GOPATH/src/github.com/matmerr/litturl/client"]
>>>>>>> 11be77e585d66ddc8be5899b213bccf00ac3893a
