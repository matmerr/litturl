FROM golang:alpine
ENV LU_DIR=$GOPATH/src/github.com/matmerr/litturl/
ADD . $LU_DIR
RUN cd $LU_DIR && \
	apk add --update git nodejs && \
	go get ./... && \
	go build main.go && \
	cd client && \
        npm install && \
	npm rebuild node-sass && \
        npm run build
EXPOSE 8001
WORKDIR $LU_DIR
CMD ["sh", "-c", "go run main.go client/dist/"]
