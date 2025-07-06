#!/bin/bash

# Lightsaber API Deployment Script
# Usage: ./deploy.sh [environment]
# Environment: dev (default) | prod

set -e

ENVIRONMENT=${1:-dev}
PROJECT_NAME="lightsaber"

echo "🚀 Deploying Lightsaber API - Environment: $ENVIRONMENT"

# Set the appropriate docker-compose file
if [ "$ENVIRONMENT" = "prod" ]; then
    COMPOSE_FILE="docker-compose.prod.yml"
    echo "📦 Using production configuration"
else
    COMPOSE_FILE="docker-compose.yml"
    echo "🛠️  Using development configuration"
fi

# Check if .env file exists
if [ ! -f .env ]; then
    echo "⚠️  Warning: .env file not found. Creating a template..."
    cat > .env << EOF
# Database
DB_PASSWORD=password

# SMTP Configuration (optional)
SMTP_HOST=sandbox.smtp.mailtrap.io
SMTP_PORT=2525
SMTP_USERNAME=
SMTP_PASSWORD=
SMTP_SENDER=

# CORS (optional)
TRUSTED_ORIGINS=*
EOF
    echo "📝 Created .env template. Please update with your values."
fi

# Pull latest images
echo "📥 Pulling latest images..."
docker-compose -f $COMPOSE_FILE pull

# Build and start services
echo "🔨 Building and starting services..."
docker-compose -f $COMPOSE_FILE up --build -d

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
sleep 10

# Check if API is responding
echo "🩺 Health check..."
if curl -f http://localhost:4000/v1/healthcheck > /dev/null 2>&1; then
    echo "✅ API is healthy and running!"
    echo "🌐 API available at: http://localhost:4000"
    echo "📊 Health check: http://localhost:4000/v1/healthcheck"
else
    echo "❌ API health check failed"
    echo "📋 Checking logs..."
    docker-compose -f $COMPOSE_FILE logs api
    exit 1
fi

# Show running containers
echo ""
echo "📋 Running containers:"
docker-compose -f $COMPOSE_FILE ps

echo ""
echo "🎉 Deployment completed successfully!"
echo ""
echo "Useful commands:"
echo "  View logs:     docker-compose -f $COMPOSE_FILE logs -f"
echo "  Stop services: docker-compose -f $COMPOSE_FILE down"
echo "  Clean up:      docker-compose -f $COMPOSE_FILE down -v"
