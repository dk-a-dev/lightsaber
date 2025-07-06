# 🚀 Lightsaber API - Complete DevOps Setup Documentation

## Project Overview
Lightsaber is a production-ready Go REST API for movie management with comprehensive DevOps practices including containerization, monitoring, CI/CD, and automated testing.

## 🛠️ Technology Stack

### Backend
- **Language:** Go 1.24
- **Framework:** Custom HTTP router with httprouter
- **Database:** PostgreSQL 16
- **Authentication:** JWT tokens
- **Logging:** Structured JSON logging

### DevOps & Infrastructure
- **Containerization:** Docker & Docker Compose
- **Database Migrations:** golang-migrate
- **Monitoring:** Grafana + Graphite
- **CI/CD:** GitHub Actions
- **Testing:** Go testing framework with race condition detection

## 📋 Prerequisites (Fresh System Setup)

### macOS Setup
```bash
# Install Homebrew (if not installed)
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install required tools
brew install docker
brew install git
brew install go

# Start Docker Desktop
open -a Docker
```

### Ubuntu/Linux Setup
```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Docker
sudo apt install docker.io docker-compose -y
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER

# Install Go
wget https://go.dev/dl/go1.24.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install Git
sudo apt install git -y
```

## 🚀 Complete Project Setup

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/lightsaber.git
cd lightsaber
```

### 2. Environment Configuration
```bash
# Copy environment template
cp .env.example .env

# Edit .env file with your configuration
# DATABASE_URL=postgres://greenlight:password@db:5432/greenlight?sslmode=disable
# TRUSTED_ORIGINS=http://localhost:3000,http://localhost:4000
```

### 3. Build and Start the Application
```bash
# Build and start all services (API, Database, Monitoring)
make docker/up

# Alternative: Manual Docker Compose
docker-compose up --build -d
```

### 4. Verify Installation
```bash
# Check if all services are running
docker-compose ps

# Test API health
curl http://localhost:4000/v1/healthcheck

# Check API metrics
curl http://localhost:4000/debug/vars
```

## 🐳 Docker Architecture

### Multi-Service Setup
```yaml
# docker-compose.yml services:
- api: Go application (Port 4000)
- db: PostgreSQL database (Port 5432) 
- migrate: Database migration service
- graphite: Metrics collection (Port 8080)
- grafana: Monitoring dashboard (Port 3000)
```

### Docker Features
- **Multi-stage builds** for optimized image size
- **Health checks** for service dependencies
- **Volume persistence** for database and metrics
- **Network isolation** between services
- **Environment-based configuration**

## 📊 Monitoring & Observability

### Metrics Collection
- **Custom Metrics Package:** Go client for Graphite
- **Middleware Integration:** Automatic request/response tracking
- **Real-time Metrics:** Request counts, response times, HTTP methods

### Grafana Dashboard
- **URL:** http://localhost:3000
- **Credentials:** admin/admin
- **Dashboards:** Pre-configured API metrics visualization

### Monitored Metrics
- Total API requests
- Requests by HTTP method (GET, POST, PATCH, DELETE)
- Response times by endpoint
- Request rate (requests per second)
- Application health status

## 🧪 Testing Strategy

### Test Coverage
```bash
# Run all tests with race detection
make test

# Run specific test suites
go test -v ./cmd/api/...
go test -v ./internal/...
```

### Test Categories
- **Unit Tests:** Data models, validators, utilities
- **Integration Tests:** API endpoints, database operations
- **Health Check Tests:** Service availability
- **Performance Tests:** Response time validation

### Generate Test Metrics
```bash
# Run comprehensive API tests
./test-metrics.sh
```

## 🔄 CI/CD Pipeline

### GitHub Actions Workflows

#### 1. Build Workflow (`.github/workflows/build.yml`)
- Triggered on push/PR to main
- Multi-Go version testing (1.23, 1.24)
- Dependency caching
- Binary compilation verification

#### 2. Test Workflow (`.github/workflows/test.yml`)
- Comprehensive test execution
- Race condition detection
- Code coverage reporting
- Multiple environment testing

### Pipeline Features
- **Parallel execution** for faster builds
- **Dependency caching** for improved performance
- **Multi-version testing** for compatibility
- **Automated security scanning**

## 🗄️ Database Management

### Migration System
```bash
# Create new migration
make db/migration/new name=add_new_feature

# Apply migrations (requires confirmation)
make db/migration/up

# Database connection
make psql
```

### Database Features
- **Automated migrations** on container startup
- **Connection pooling** for performance
- **Health checks** for reliability
- **Data persistence** with Docker volumes

## 🔧 Development Workflow

### Local Development
```bash
# Start development environment
make docker/up

# View logs
make docker/logs

# API-specific logs
make docker/logs/api

# Stop services
make docker/down
```

### Production Deployment
```bash
# Production build
make docker/prod/up

# Deploy script
./deploy.sh prod
```

## 🛡️ Security Features

### Implementation
- **JWT Authentication** with secure token generation
- **Rate limiting** to prevent abuse
- **CORS configuration** for web security
- **Input validation** on all endpoints
- **SQL injection protection** with parameterized queries

### Environment Security
- **Environment variables** for sensitive data
- **Docker secrets** for production credentials
- **.env.example** for safe sharing
- **Gitignore** protection for credentials

## 📈 Performance Optimization

### Application Level
- **Connection pooling** for database efficiency
- **Middleware optimization** for request processing
- **Structured logging** for performance monitoring
- **Graceful shutdown** for reliability

### Infrastructure Level
- **Docker multi-stage builds** for smaller images
- **Alpine Linux base** for security and size
- **Health checks** for automatic recovery
- **Resource limits** for container management

## 🔍 Monitoring Setup Instructions

### Step-by-Step Dashboard Access

1. **Start the monitoring stack:**
   ```bash
   make docker/up
   ```

2. **Access Grafana:**
   - URL: http://localhost:3000
   - Username: admin
   - Password: admin

3. **Generate metrics:**
   ```bash
   ./test-metrics.sh
   ```

4. **View dashboard:**
   - Navigate to "Lightsaber API Dashboard"
   - Observe real-time metrics
   - Monitor API performance

### Dashboard Panels
- **Total Requests:** Counter of all API calls
- **HTTP Methods:** Breakdown by GET/POST/PATCH/DELETE
- **Response Times:** Performance metrics by endpoint
- **Request Rate:** Real-time throughput monitoring

## 🚀 Deployment Options

### Local Development
```bash
make docker/up      # Development with hot reload
make docker/logs    # Monitor application logs
```

### Production Environment
```bash
make docker/prod/up   # Production-optimized containers
./deploy.sh prod      # Automated production deployment
```

### Cloud Deployment Ready
- **Environment-based configuration**
- **Health checks** for load balancers
- **Graceful shutdown** for zero-downtime deploys
- **Metrics endpoints** for monitoring systems

## 🏗️ Architecture Benefits

### Microservices Ready
- **Service separation** with Docker Compose
- **Independent scaling** capabilities
- **Service discovery** through Docker networks
- **Health check dependencies**

### DevOps Best Practices
- **Infrastructure as Code** with Docker Compose
- **Automated testing** with GitHub Actions
- **Monitoring and alerting** with Grafana
- **Database migrations** with version control

### Development Experience
- **One-command setup** with Make targets
- **Hot reload** for development
- **Comprehensive logging** for debugging
- **API documentation** through code

## 📝 Available Make Commands

```bash
make help              # Show all available commands
make run               # Run API locally
make test              # Execute all tests
make audit             # Code quality checks
make build             # Build binaries
make docker/build      # Build Docker image
make docker/up         # Start all services
make docker/down       # Stop all services
make docker/prod/up    # Production deployment
make deploy            # Deploy to development
make deploy/prod       # Deploy to production
```

## 🎯 Project Achievements

### Technical Excellence
✅ **Production-ready API** with comprehensive error handling  
✅ **Complete containerization** with Docker & Docker Compose  
✅ **Real-time monitoring** with Grafana & Graphite  
✅ **Automated CI/CD** with GitHub Actions  
✅ **Database migrations** with version control  
✅ **Comprehensive testing** with race condition detection  
✅ **Security implementation** with JWT and rate limiting  
✅ **Performance monitoring** with custom metrics  

### DevOps Implementation
✅ **Infrastructure as Code** approach  
✅ **Automated deployment** pipelines  
✅ **Health checking** and service dependencies  
✅ **Monitoring and observability** stack  
✅ **Configuration management** with environment variables  
✅ **Secret management** best practices  
✅ **Multi-environment support** (dev/prod)  
✅ **One-command deployment** automation  

## 🔗 Service URLs

| Service | URL | Purpose |
|---------|-----|---------|
| API | http://localhost:4000 | Main application |
| Health Check | http://localhost:4000/v1/healthcheck | Service status |
| API Metrics | http://localhost:4000/debug/vars | Application metrics |
| Grafana | http://localhost:3000 | Monitoring dashboard |
| Graphite | http://localhost:8080 | Raw metrics interface |
| PostgreSQL | localhost:5432 | Database connection |

---

## 📞 Support & Troubleshooting

### Common Issues

1. **Port conflicts:** Ensure ports 3000, 4000, 5432, 8080 are available
2. **Docker not running:** Start Docker Desktop/daemon
3. **Permission issues:** Add user to docker group (Linux)
4. **Migration failures:** Check database connectivity

### Debug Commands
```bash
# Check service status
docker-compose ps

# View service logs
docker-compose logs [service-name]

# Access database
make psql

# Test API connectivity
curl -v http://localhost:4000/v1/healthcheck
```

This comprehensive setup demonstrates production-ready DevOps practices with modern tooling and monitoring capabilities, suitable for enterprise-level applications.
