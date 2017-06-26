FROM golang:latest
RUN mkdir /src
RUN go get github.com/matmerr/litturl
RUN cd client && \
    npm install && \
    npm run build
CMD ["go", "run", "main.go"] 