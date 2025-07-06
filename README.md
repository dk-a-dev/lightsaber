# Lightsaber API

A modern Go REST API with Docker deployment.

## 🚀 Quick Start

### Development
```bash
# Start development environment
make docker/up

# Run tests
make test

# View logs
make docker/logs/api
```

### Production Deployment
```bash
# Deploy to production
make deploy/prod

# Or use the deployment script directly
./deploy.sh prod
```

## 📋 Available Commands

### Development
- `make run` - Run API locally
- `make test` - Run all tests
- `make docker/up` - Start dev services
- `make docker/down` - Stop dev services

### Production
- `make deploy/prod` - Deploy to production
- `make docker/prod/up` - Start production services
- `make docker/prod/down` - Stop production services

### Logs & Monitoring
- `make docker/logs` - View all logs
- `make docker/logs/api` - View API logs only

## 🔧 Configuration

Create a `.env` file:
```env
# Database
DB_PASSWORD=your_secure_password

# SMTP (optional)
SMTP_HOST=sandbox.smtp.mailtrap.io
SMTP_PORT=2525
SMTP_USERNAME=your_username
SMTP_PASSWORD=your_password
SMTP_SENDER=your@email.com

# CORS (optional)
TRUSTED_ORIGINS=https://yourdomain.com
```

## 🌐 API Endpoints

- **Health Check**: `GET /v1/healthcheck`
- **Movies**: `GET|POST /v1/movies`
- **Users**: Authentication endpoints
- **Tokens**: Token management

## 🏗️ Architecture

- **Go 1.24** - Modern Go backend
- **PostgreSQL 16** - Database with migrations
- **Docker** - Containerized deployment
- **GitHub Actions** - CI/CD pipeline
- **Comprehensive Testing** - Unit and integration tests

## 📊 Production Features

- ✅ Health checks and monitoring
- ✅ Graceful shutdowns
- ✅ Database migrations
- ✅ Structured logging
- ✅ Security best practices
- ✅ Automatic restarts

## 🔒 Security

- Non-root container user
- Environment variable secrets
- CORS configuration
- Rate limiting
- Input validation

## 📈 Monitoring

Access your API:
- **Development**: http://localhost:4000
- **Production**: http://localhost (port 80)
- **Health Check**: `/v1/healthcheck`

---

Built with ❤️ using Go and Docker
