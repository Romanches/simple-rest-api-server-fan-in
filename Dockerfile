FROM alpine:latest

# Set some configurations
ENV DB_HOST="localhost"

RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY go-app /app/
#COPY ./cmd/config.yaml /app/

LABEL Name=go-app

# Expose your port
EXPOSE 8080

# Indicate the binary as our entrypoint
ENTRYPOINT "/app/go-app"

#CMD ["/app/go-app"]