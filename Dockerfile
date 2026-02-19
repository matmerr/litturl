FROM golang:1.24-alpine AS go-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --update git && go mod download
COPY server server
COPY main.go .
RUN go build -o /app/litturl .

FROM node:22-alpine AS node-stage
WORKDIR /app
COPY client/package.json client/package-lock.json ./
RUN npm ci
COPY client/ .
RUN npm run build

FROM alpine:latest
WORKDIR /app
COPY conf conf
COPY --from=go-stage /app/litturl .
COPY --from=node-stage /app/dist/ dist/
EXPOSE 8001
CMD ["./litturl", "dist/"]

