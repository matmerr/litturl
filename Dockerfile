FROM golang:alpine as go-stage
ENV LU_DIR=$GOPATH/src/github.com/matmerr/litturl/
WORKDIR $LU_DIR
ADD server server
ADD main.go .
RUN apk add --update git && go get ./...
RUN go build  -o /app/litturl main.go

FROM node:alpine as node-stage
RUN apk add --update git 
WORKDIR /app
COPY client .
RUN npm install && npm run build

from alpine:latest
WORKDIR /app
COPY --from=go-stage /app/litturl .
COPY --from=node-stage /app/dist/ dist/
CMD ["./litturl", "dist/"]





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



