FROM node:latest AS frontend
WORKDIR /app
COPY . /app
RUN npm install && npm run prod

FROM golang:1.20-alpine AS backend
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o tree .

FROM alpine:latest
COPY . /app
WORKDIR /app/static
COPY --from=frontend /app/static/ .
WORKDIR /app
COPY --from=backend /app/tree .
EXPOSE 8000
CMD ["./tree"]