# # Use an official Golang runtime as a parent image
# FROM golang:latest

# # Set the working directory to /app
# WORKDIR /app

# # Copy the current directory contents into the container at /app
# COPY . /app

# # Install Node.js
# RUN apt-get update && apt-get install -y nodejs npm

# # Install the Node.js packages
# RUN npm install

# # Build css assets
# RUN npm run prod

# # Install go dependencies
# RUN go mod tidy

# # Build the Go application
# RUN go build -o tree .

# # Expose port 8000 for the Go application
# EXPOSE 8000

# # Run the Go application
# CMD ["./tree"]

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