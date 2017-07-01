FROM golang:latest
RUN mkdir /app && cd /app
RUN go get github.com/matmerr/litturl
RUN cd client && \
    npm install && \
    npm run build && \
    cd /app

CMD ["go", "run", "main.go"] 