# Build Stage
FROM node:20-alpine AS node-builder
WORKDIR /app
# Copy all files needed for CSS processing
COPY package*.json ./
COPY tailwind.config.js ./
COPY static/css ./static/css
COPY templates ./templates
# Install dependencies and build CSS
RUN npm install
RUN npm run css:build

# Go Builder Stage
FROM golang:1.24-alpine AS go-builder
WORKDIR /app
# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest
# Copy everything
COPY . .
# Copy the built CSS from node stage
COPY --from=node-builder /app/static/css/main.css ./static/css/
# Generate templ files and build
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./cmd/server/main.go

# Final Stage
FROM alpine:latest
WORKDIR /app
# Copy the binary
COPY --from=go-builder /app/server .
# Copy static files including CSS
COPY --from=node-builder /app/static ./static
# Copy templates for reference
COPY --from=go-builder /app/templates ./templates
# Expose port
EXPOSE 8080
# Run the binary
CMD ["./server"]
