# Backend Stage
FROM golang:1.20 AS backend-builder
WORKDIR /app/backend

# Copy Go files and modules
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN go build -o backend-app main.go

# Frontend Stage
FROM node:18 AS frontend-builder
WORKDIR /app/frontend

# Copy package.json and install dependencies
COPY frontend/package*.json ./
RUN npm install

# Copy the rest of the frontend code
COPY frontend/ ./
RUN npm run build

# Final Stage - Combine Backend and Frontend
FROM alpine:3.18
WORKDIR /app

# Install required dependencies for the backend
RUN apk add --no-cache ca-certificates

# Copy built backend binary
COPY --from=backend-builder /app/backend/backend-app /app/backend-app

# Copy frontend build
COPY --from=frontend-builder /app/frontend/build /app/frontend/build

# Expose port for the application
EXPOSE 8080

# Command to run the backend
CMD ["./backend-app"]
