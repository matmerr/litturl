FROM golang:alpine
ADD . $LU_DIR
WORKDIR /go/src/github.com/matmerr/litturl/
COPY server/ main.go ./
RUN go build -a -installsuffix -o litturl main.go

FROM node:alpine
RUN apk add --update git 
COPY client ./
RUN cd client && npm install && npm run build
COPY --from=0 /go/src/github.com/matmerr/litturl/ .
CMD ["./litturl client/dist/"]







#RUN cd $LU_DIR && \
#	apk add --update git nodejs && \
#	go get ./... && \
#	go build main.go && \
#	cd client && \
#        npm install && \
#	#npm rebuild node-sass && \
#        npm run build
#EXPOSE 8001
#WORKDIR $LU_DIR



