FROM golang:alpine

RUN apk update && apk add --no-cache git

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .
RUN go mod tidy

# Set environment variables
ENV ENVIRONMENT production

# Build
RUN go build -o ./bin/app ./cmd

# Optional:
EXPOSE 8080

# Run
CMD [ "./bin/app" ]