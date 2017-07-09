FROM golang:latest
RUN curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash - && \
    sudo apt-get install -y nodejs
RUN mkdir /app && cd /app
RUN go get github.com/matmerr/litturl
RUN cd client && \
    npm install && \
    npm run build && \
    cd /app
CMD ["go", "run", "main.go"] 